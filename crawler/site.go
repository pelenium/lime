package main

import "github.com/securisec/go-keywords"

type site struct {
	url      string
	title    string
	htmlCode string
	keywords []string
}

func (s *site) SetKW(data string) {
	kw, _ := keywords.Extract(data, keywords.ExtractOptions{
		StripTags:        true,
		RemoveDuplicates: true,
		IgnorePattern:    "<.+>",
		Lowercase:        true,
		AddStopwords:     []string{".", "..", "...", ",", "!", "?", "*"},
	})

	for _, i := range kw {
		s.keywords = append(s.keywords, i)
	}
}
