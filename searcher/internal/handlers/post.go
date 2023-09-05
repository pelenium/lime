package handlers

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func Post(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	data := string(body)

	query := gjson.Get(data, "query").String()
	words := strings.Split(query, " ")

	fmt.Println(words)
	fmt.Println(query)
}
