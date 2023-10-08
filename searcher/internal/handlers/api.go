package handlers

import (
	"net/http"
	"searcher/internal/engine"
	"strings"

	"github.com/gin-gonic/gin"
)

func Api(c *gin.Context) {
	req := strings.ToLower(c.Params.ByName("req"))
	query := strings.Split(req, "_")
	sites := engine.GetLinks(query)

	var jsn []map[string]interface{}
	for _, site := range sites {
		jsn = append(jsn, map[string]interface{}{
			"url":      site.Url,
			"keywords": site.Keywords,
			"title":    site.Title,
			"rating":   site.Rating,
		})
	}

	c.JSON(http.StatusOK, jsn)
}
