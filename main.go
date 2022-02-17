package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/emlog/goexample/api"
)

// @title 简易笔记工具
// @version 1.0
// @description 简易笔记发布系统 - GO学习参考项目
// @termsOfService https://github.com/emlog/goexample
func main() {

	// 创建监听退出chan
	c := make(chan os.Signal)
	// 监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
			case syscall.SIGUSR1:
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	router := gin.Default()
	api.Router(router)
	router.Run(":5956")
}
