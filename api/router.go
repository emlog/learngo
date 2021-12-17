package api

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	note := router.Group("/note")
	{
		note.POST("/create", NoteCreate)
		// note.POST("/read", readEndpoint)
		// note.POST("/update", updateEndpoint)
		// note.POST("/delete", deleteEndpoint)
	}

}
