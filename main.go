package main

import (
	"github.com/gin-gonic/gin"

	"github.com/emlog/goexample/api"
)

// @title 简易笔记工具
// @version 1.0
// @description 简易笔记发布系统 - GO学习参考项目
// @termsOfService https://github.com/emlog/goexample
func main() {
	router := gin.Default()
	api.Router(router)
	router.Run(":8080")
}
