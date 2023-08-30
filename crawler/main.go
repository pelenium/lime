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
	"golang.org/x/net/html"
)

var (
	links  = make([]string, 0)
	used   = make([]string, 0)
	insert = "INSERT INTO sites (url, title, keywords, html) values ($1, $2, $3, $4)"
)

func main() {
	conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	links = append(links, "https://vc.ru/dev/225726-sayty-na-kotoryh-mozhno-ottochit-navyki-resheniya-zadach-po-programmirovaniyu")

	for len(links) != 0 {
		fmt.Println(links[0])
		goToPage(db, links[0])
		fmt.Println("it's ok")
		fmt.Println()

		links = links[1:]
	}
}

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
					links = append(links, strings.Split(strings.Split(link, "\n")[0], " ")[0])
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
	s.SetKW(string(data))

	_, err = db.Exec(insert, s.url, s.title, pq.Array(s.keywords), s.htmlCode)
	if err != nil {
		fmt.Println(err)
		return
	}
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
