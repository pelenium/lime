package engine

import (
	"database/sql"
	"fmt"
)

func findInKeywords(str string) []site {
	conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result []site

	req := `SELECT * FROM sites WHERE EXISTS (SELECT 1 FROM unnest(keywords) AS element WHERE lower(element::text) LIKE lower('%$1%'))`

	rows, err := db.Query(req, str)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		s := site{}
		err = rows.Scan(&s.url, &s.title, &s.keywords, &s.htmlCode, &s.rating)
		if err != nil {
			fmt.Println(err)
			continue
		}
		result = append(result, s)
	}

	return result
}
