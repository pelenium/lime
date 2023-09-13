package main

import (
	"database/sql"
	"fmt"
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

	links = append(links, "https://pikabu.ru/")

	for len(links) != 0 {
		fmt.Println(links[0])
		goToPage(db, links[0])
		fmt.Println("it's ok")
		fmt.Println()

		links = links[1:]
	}
}
