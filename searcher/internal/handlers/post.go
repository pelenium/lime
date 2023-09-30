package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func IndexPost(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	data := string(body)

	query := strings.Join(strings.Split(gjson.Get(data, "query").String(), " "), "_")

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/%s", query))
}
