package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	_ "database/sql"

	"github.com/antchfx/htmlquery"
	_ "github.com/lib/pq"
	"github.com/securisec/go-keywords"
)

func main() {
	//conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	//db, err := sql.Open("postgres", conn)
	//if err != nil {
	//	panic(err)
	//}

	goToPage("https://habr.com/ru/companies/skbkontur/articles/723840/")
}

func goToPage(url string) {
	req, _ := http.Get(url)
	data, _ := io.ReadAll(req.Body)
	k, _ := keywords.Extract(string(data), keywords.ExtractOptions{
		StripTags:        true,
		RemoveDuplicates: true,
		IgnorePattern:    "<.+>",
		Lowercase:        true,
		AddStopwords:     []string{".", "..", "...", ",", "!", "?", "*"},
	})

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		panic(err)
	}

	list := htmlquery.Find(doc, "//a/@href")

	urls := make([]string, 0)

	for _, n := range list {
		content := htmlquery.SelectAttr(n, "href")
		if strings.Contains(content, "://") {
			urls = append(urls, content)
		}
	}

	fmt.Println(k)
	fmt.Println()
	fmt.Println(strings.Join(urls, "\n"))
}
