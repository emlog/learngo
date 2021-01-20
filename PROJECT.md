# Go语言示例和学习笔记
这里包含了Go语言的学习笔记和一个完整可以运行的示例，方便快速了解go语言的基本语法和特性。

# 项目示例

## 项目目录结构

* configs 配置文件
* docs 文档
* example 演示代码示例（和项目无关）
* global 全局变量
* internal 内部模块
    * dao 数据访问层（database access object）所有数据相关操作都在dao层，包括MySQL，ES等
    * middleware HTTP中间件
    * model 模型层，用于存放model对象
    * routers 路由相关
    * service 项目核心业务逻辑
* pkg 项目相关的模块包
* script 各类构建、安装、分析等操作的脚本
* storage 项目生成的临时文件
* third_party 第三方资源工具，如 Swagger UI

## 项目准备

初始化项目
$ go mod init github.com/emlog/goexample

安装gin
$ go get -u github.com/gin-gonic/gin

goimports 和 goreturns 是格式化和包引用管理，建议使用 goreturns
$ go get golang.org/x/tools/cmd/goimports
