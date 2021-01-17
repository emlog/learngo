/*
函数示例
*/
package example

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func init() {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}
