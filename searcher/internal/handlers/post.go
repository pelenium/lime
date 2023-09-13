package handlers

import (
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
	engine.GetLinks(query)
}
