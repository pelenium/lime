package engine

import (
	"database/sql"
	"fmt"
)

func findByUrl(w string) []site {
	conn := "user=postgres password=781842 dbname=lime sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT (url, title) FROM sites WHERE url=$1", w)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var sites []site

	for rows.Next() {
		s := site{}
		if err := rows.Scan(&s.url, &s.title); err != nil {
			fmt.Println(err)
			continue
		}
		s.rating += 0.15
		sites = append(sites, s)
	}
	return sites
}
