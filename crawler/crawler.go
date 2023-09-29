package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/net/html"
)

func goToPage(db *sql.DB, url string) {
	used = append(used, url)
	sort.SliceStable(used, func(i, j int) bool {
		return used[i] < used[j]
	})

	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return
	}

	list := htmlquery.Find(doc, "//a/@href")

	for _, n := range list {
		link := htmlquery.SelectAttr(n, "href")
		if len(link) >= 4 {
			if link[:4] == "http" {
				if !contains(used, strings.Split(strings.Split(link, "\n")[0], " ")[0]) {
					_, err := db.Exec(updateCounter, strings.Split(strings.Split(link, "\n")[0], " ")[0])
					if err != nil {
						fmt.Println(err)
					}
					links = append(links, strings.Split(strings.Split(link, "\n")[0], " ")[0])
				} else {
					fmt.Println("contains")
					return
				}
			}
		}
	}

	title := getTitle(doc)

	s := site{
		url:      url,
		title:    title,
		htmlCode: string(data),
	}
	s.setKW(string(data))

	_, err = db.Exec(insert, s.url, s.title, pq.Array(s.keywords), s.htmlCode, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("it's ok")
	fmt.Println()
}

func contains(arr []string, e string) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := (left + right) / 2

		if arr[middle] == e {
			return true
		} else if arr[middle] < e {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return false
}

func getTitle(doc *html.Node) (result string) {
	t := htmlquery.Find(doc, "//head/title/text()")
	for _, n := range t {
		result = n.Data
	}
	fmt.Println(result)
	return
}
