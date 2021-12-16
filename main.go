package main

import (
	"github.com/emlog/goexample/api"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api.Router(router)

	router.Run(":8080")

}
