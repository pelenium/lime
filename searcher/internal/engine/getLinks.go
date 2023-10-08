package engine

import (
	"sort"
	"strings"
)

func GetLinks(query []string) []Site {
	var result []Site

	for _, i := range query {
		for _, link := range findByUrl(i) {
			if indx, c := contains(result, link); !c {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url < result[j].Url
				})
			} else {
				result[indx].Rating += link.Rating
			}
		}

		for _, link := range findByTitle(i) {
			if indx, c := contains(result, link); !c {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url < result[j].Url
				})
			} else {
				result[indx].Rating += link.Rating
			}
		}

		for _, link := range findInKeywords(i) {
			if indx, c := contains(result, link); !c {
				result = append(result, link)
				sort.Slice(result, func(i, j int) bool {
					return result[i].Url < result[j].Url
				})
			} else {
				result[indx].Rating += link.Rating
			}
		}
	}

	for _, link := range findByUrl(strings.Join(query, " ")) {
		if indx, c := contains(result, link); !c {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url < result[j].Url
			})
		} else {
			result[indx].Rating += link.Rating
		}
	}

	for _, link := range findByTitle(strings.Join(query, " ")) {
		if indx, c := contains(result, link); !c {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url < result[j].Url
			})
		} else {
			result[indx].Rating += link.Rating
		}
	}

	for _, link := range findInKeywords(strings.Join(query, " ")) {
		if indx, c := contains(result, link); !c {
			result = append(result, link)
			sort.Slice(result, func(i, j int) bool {
				return result[i].Url < result[j].Url
			})
		} else {
			result[indx].Rating += link.Rating
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Rating > result[j].Rating
	})

	return result
}

func contains(arr []Site, e Site) (int, bool) {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := (left + right) / 2
		if arr[middle].Url == e.Url {
			return middle, true
		} else if arr[middle].Url < e.Url {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1, false
}
