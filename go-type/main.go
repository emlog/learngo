package main

import "fmt"

//定义全局变量
var x = 100

const y = 123

func main() {
	//println 是 builtin 包提供，语言内置，而 fmt.Println 来自标准库
	// The println built-in function formats its arguments in an
	// implementation-specific way and writes the result to standard error.
	// Spaces are always added between arguments and a newline is appended.
	// Println is useful for bootstrapping and debugging; it is not guaranteed
	// to stay in the language.
	println(&x, x)

	//简单模式 定义局部变量
	x := "abc"
	fmt.Println(&x, x)

	println(y)

}
