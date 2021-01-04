package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Hello, world.\n")

	//类型转换，int to string
	a := 65
	b := strconv.Itoa(a)
	fmt.Printf("b is %s \n", b)

}
