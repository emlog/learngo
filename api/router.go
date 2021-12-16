package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	// Simple group: v1
	// v1 := router.Group("/v1")
	// {
	// 	v1.POST("/login", loginEndpoint)
	// 	v1.POST("/submit", submitEndpoint)
	// 	v1.POST("/read", readEndpoint)
	// }

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run(":8080")

}
