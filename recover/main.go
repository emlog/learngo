// recover
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
