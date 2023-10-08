package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", gin.H{})
}
