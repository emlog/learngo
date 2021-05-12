/*
panic处理机制
*/
package error_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// 恢复一个panic
func TestPanicRecover(t *testing.T) {
	defer func() {
		// 捕获一个异常，程序没有终止
		if err := recover(); err != nil {
			fmt.Println("recover from： ", err)
		}
	}()

	fmt.Println("start")
	panic(errors.New("someting wrong"))

	os.Exit(-1)

}

func TestPanic(t *testing.T) {

	defer func() {
		fmt.Println("Finally")
	}()

	fmt.Println("start")

	// 会执行defer函数
	panic(errors.New("someting wrong"))

	// 不会执行defer函数
	//os.Exit(-1)

}
