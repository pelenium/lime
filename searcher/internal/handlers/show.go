package handlers

import (
	"net/http"
	"searcher/internal/engine"
	"strings"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	query := strings.Split(c.Param("req"), "_")
	sites := engine.GetLinks(query)

	jsn := map[string]interface{}{}
	for _, site := range sites {
		jsn[site.Url] = map[string]interface{}{
			"keywords": site.Keywords,
			"title":    site.Title,
		}
	}

	c.JSON(http.StatusOK, jsn)
}
