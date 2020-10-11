package main

import "fmt"

/*
定义数组
*/
var colorsList [3]string
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//忽略长度的定义
var balance2 = [...]string{"aaa", "bbb", "ccc"}

//定义切片
var slice1 []string

func main() {
	animals := [3]string{"cat", "dog", "monkey"}

	for i := 0; i <= 2; i++ {
		fmt.Printf("animal is %s \n", animals[i])
	}

	slice2 := []string{"red", "blue", "green"}

	fmt.Printf("", slice2[0])

}
