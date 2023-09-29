package main

import (
	"database/sql"
	"fmt"
)

var (
	links         = make([]string, 0)
	used          = make([]string, 0)
	insert        = "INSERT INTO sites (url, title, keywords, html, counter) values ($1, $2, $3, $4, $5)"
	updateCounter = "UPDATE sites SET counter = counter + 0.1 WHERE url = $1"
)

func main() {
	conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	links = append(links, "https://go.dev/")
	links = append(links, "https://habr.com/ru/articles/731046/")
	links = append(links, "https://blog.skillfactory.ru/glossary/golang/")
	links = append(links, "https://metanit.com/go/tutorial/")
	links = append(links, "https://habr.com/ru/hubs/programming/articles/")

	for len(links) != 0 {
		fmt.Println(links[0])
		goToPage(db, links[0])
		links = links[1:]
	}
}
