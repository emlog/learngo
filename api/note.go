package api

import (
	"github.com/gin-gonic/gin"

	"github.com/emlog/goexample/model/request"
	"github.com/emlog/goexample/model/respone"
	"github.com/emlog/goexample/service"
)

// NoteCreate 创建笔记
// @Summary 写笔记
// @Description 写入一条笔记
// @Tags Note
// @Accept json
// @Produce json
// @Param object body request.ReqNote true "请求参数"
// @Success 200 {object} respone.RespNote
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
	c.JSON(200, respone.RespNote{
		Message: "success",
		Id:      id,
	})
}
