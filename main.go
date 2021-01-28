package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// main函数只能用于main包中，且只能定义一个。
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("starting web server")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
