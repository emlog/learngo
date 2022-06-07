// recover
// Goroutine能够实现高并发，但是如果某个Goroutine panic了，而且这个Goroutine里面没有捕获recover，那么整个进程就会挂掉。所以，好的习惯是每当go产生一个goroutine，就需要写下recover。
package main

import "fmt"

func main() {
	a := []string{"a", "b"}
	go checkAndPrintWithRecover(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	checkAndPrint(a, 2)
}

func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

// recover function in the same goroutine as panic
func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
