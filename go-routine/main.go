package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	go say("world")
	say("hello")

	//for channel
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:3], c)
	go sum(s[3:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y)

}
