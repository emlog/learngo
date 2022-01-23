package main

import (
	"github.com/gin-gonic/gin"

	"github.com/emlog/goexample/api"
)

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	router := gin.Default()
	api.Router(router)
	router.Run(":8080")
}
