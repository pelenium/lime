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
	r.GET("/search/:req", handlers.Show)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
