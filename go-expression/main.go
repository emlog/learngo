/*
表达式示例：控制结构
*/
package main

import "fmt"

//var a = "hello world"

func main() {
	fmt.Println("hello world!!")

	//定义切片
	s := []int{1, 2, 3}

	//遍历切片
	for _, value := range s {
		fmt.Printf("%d , %s ", value, "---")
	}

}
