package engine

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func findInKeywords(str string) []Site {
	conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var result []Site
	req := `SELECT * FROM sites WHERE EXISTS (SELECT 1 FROM unnest(Keywords) AS element WHERE lower(element::text) LIKE lower('%` + str + `%'))`

	rows, err := db.Query(req)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		s := Site{}
		err = rows.Scan(&s.Url, &s.Title, pq.Array(&s.Keywords), &s.HtmlCode)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.Rating = 1.2
		result = append(result, s)
	}

	return result
}
