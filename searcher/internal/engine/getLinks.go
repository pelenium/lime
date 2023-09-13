package engine

func GetLinks(query []string) (result string) {
	var sites []site
	for _, i := range query {
		for _, link := range findByUrl(i) {
			sites = append(sites, link)
		}
	}
	return
}
