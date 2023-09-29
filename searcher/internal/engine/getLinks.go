package engine

import "sort"

func GetLinks(query []string) []site {
	var result []site

	for _, i := range query {
		for _, link := range findByUrl(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].url > result[j].url
				})
			}
		}

		for _, link := range findByTitle(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].url > result[j].url
				})
			}
		}

		for _, link := range findInKeywords(i) {
			if !contains(result, link) {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].url > result[j].url
				})
			}
		}
	}

	return result
}

func contains(arr []site, el site) bool {
	left, right := -1, len(arr)

	for left < right-1 {
		middle := (left + right) / 2

		if arr[middle].url == el.url {
			return true
		} else if arr[middle].url < el.url {
			left = middle
		} else {
			right = middle
		}
	}

	return false
}
