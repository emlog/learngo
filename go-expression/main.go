/*
表达式示例：控制结构
*/
package main

import (
	"fmt"
)

//var a = "hello world"

func main() {
	fmt.Println("hello world!!")

	//if
	x := 5
	if x > 0 {

	}

	//for: 遍历切片
	s := []int{1, 2, 3}
	for _, value := range s {
		fmt.Printf("%d , %s ", value, "---")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("@")
	}

	// switch ,若 switch 后面没有表达式，它将匹配 true
	c := 5
	switch {
	case c > 6:
		fmt.Println("1111")
	case c < 3:
		fmt.Println("2222")
	case c == 5:
		fmt.Println("3333")
	}

}
