package engine

import (
	"sort"
	"strings"
)

func GetLinks(query []string) []Site {
	var result []Site

	for _, i := range query {
		for _, link := range findByUrl(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url > result[j].Url
				})
			}
		}

		for _, link := range findByTitle(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url > result[j].Url
				})
			}
		}

		for _, link := range findInKeywords(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url > result[j].Url
				})
			}
		}
	}

	for _, link := range findByUrl(strings.Join(query, " ")) {
		if !contains(result, link) {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url > result[j].Url
			})
		}
	}

	for _, link := range findByTitle(strings.Join(query, " ")) {
		if !contains(result, link) {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url > result[j].Url
			})
		}
	}

	for _, link := range findInKeywords(strings.Join(query, " ")) {
		if !contains(result, link) {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url > result[j].Url
			})
		}
	}

	return result
}

func contains(arr []Site, el Site) bool {
	left, right := -1, len(arr)

	for left < right-1 {
		middle := (left + right) / 2

		if arr[middle].Url == el.Url {
			return true
		} else if arr[middle].Url < el.Url {
			left = middle
		} else {
			right = middle
		}
	}

	return false
}
