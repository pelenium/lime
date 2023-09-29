package handlers

import (
	"encoding/json"
	"io"
	"searcher/internal/engine"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func IndexPost(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	data := string(body)

	query := strings.Split(gjson.Get(data, "query").String(), " ")
	sites := engine.GetLinks(query)

	jsn := map[string]interface{}{}
	for _, site := range sites {
		jsn[site.Url] = map[string]interface{}{
			"keywords": site.Keywords,
			"title":    site.Title,
		}
	}

	// json with sites
	jsnResult, err := json.Marshal(jsn)
	if err != nil {
		panic(err)
	}

}
