package api

import (
	"github.com/gin-gonic/gin"

	_ "github.com/emlog/goexample/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router(router *gin.Engine) {

	note := router.Group("/note")
	{
		note.POST("/create", NoteCreate)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
