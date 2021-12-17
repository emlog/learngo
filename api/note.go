package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NoteCreate(c *gin.Context) {
	message := c.PostForm("message")

	fmt.Printf("message: %s", message)

	c.JSON(200, gin.H{
		"message": message,
	})
}
