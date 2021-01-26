/*
代码示例：初始化
*/
package main

import "fmt"

// int和int64运行的结果一样。int64是有符号 64 位整型，而在32位操作系统中int的大小也是32位（4字节）。
var x int32 = 100
var b1 *int
var b2 []int
var b3 map[string]int
var b4 chan int
var b5 func(string) int
var b6 error // error 是接口

// 定义常量
const y = 123
const (
	number = 100
	name   = "linda"
)
const (
	a = iota
	b
	c
	d
	e
)

// main 初始化相关的例子
func main() {
	// println 是 builtin 包提供，语言内置，而 fmt.Println 来自标准库
	// The println built-in function formats its arguments in an
	// implementation-specific way and writes the result to standard error.
	// Spaces are always added between arguments and a newline is appended.
	// Println is useful for bootstrapping and debugging; it is not guaranteed
	// to stay in the language.
	println(&x, x)

	// 简单模式 定义局部变量
	x := "abc"
	fmt.Println(&x, x)

	println(y) // print 123

	fmt.Println(e) // print 4

}
