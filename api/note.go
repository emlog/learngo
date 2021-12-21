package api

import (
	"github.com/emlog/goexample/model/request"
	"github.com/emlog/goexample/service"
	"github.com/gin-gonic/gin"
)

func NoteCreate(c *gin.Context) {
	req := &request.ReqNote{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.NewNoteService().CreateNote(req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"id":      id,
	})
}
