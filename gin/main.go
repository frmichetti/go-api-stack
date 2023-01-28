package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(static.Serve("/", static.LocalFile("./public/index.html", false)))
	r.Static("/public", "./public")
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("more_static"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
