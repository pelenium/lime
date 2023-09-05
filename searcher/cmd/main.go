package main

import (
	"searcher/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/internal/static", "./internal/static")
	r.LoadHTMLGlob("./internal/static/html/*.html")

	r.GET("/", handlers.Index)
	r.POST("/api/:req", handlers.Post)
	r.GET("/:req", handlers.Show)

	r.Run(":8080")
}
