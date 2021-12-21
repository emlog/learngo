package api

import (
	"github.com/emlog/goexample/model/request"
	"github.com/emlog/goexample/service"
	"github.com/gin-gonic/gin"
)

// NoteCreate 创建笔记
// @Summary 写笔记
// @Description 写入一条笔记
// @Tags Note
// @Accept json
// @Produce json
// @Param object body request.ReqNote{} true "请求参数"
// @Success 200 {object} interface{} "返回结果"
// @Router /note/create [post]
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
