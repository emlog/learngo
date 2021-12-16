package server

import (
	"encoding/json"
	"net/http"

	"github.com/emlog/goexample/store"
	"github.com/gin-gonic/gin"
)

type BookStoreServer struct {
	s store.Store
}

func (bs *BookStoreServer) CreateBookHandler(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func (bs *BookStoreServer) GetBookHandler(c *gin.Context) {
	id := c.Param("name")

	book, err := bs.s.Get(id)
	if err != nil {
		return
	}

	response(c, book)
}

func response(c *gin.Context, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, data)
}
