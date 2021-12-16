package main

import (
	_ "github.com/emlog/goexample/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/book/detail", BookStoreServer.GetBookHandler)
	router.POST("/book/add", BookStoreServer.CreateBookHandler)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
