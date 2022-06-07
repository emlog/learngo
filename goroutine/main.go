package main

import (
	"fmt"
	"time"
)

func test() {
	go func() {
		fmt.Println("hello a")
		return
	}()

	go func() {
		fmt.Println("hi b")
	}()

	fmt.Println("hi c")
}

func main() {
	test()
	time.Sleep(time.Second * 1)
}
