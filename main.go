package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Books struct {
	ID    int    `json:"id"`
	Nama  string `json:"name"`
	Title string `json:"title"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
