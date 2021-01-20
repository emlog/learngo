package api

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Create(c *gin.Context) {

}
