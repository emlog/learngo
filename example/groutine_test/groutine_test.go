package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i) // 输出顺序可能不是1234567…… 是因为协程的调用并不是顺序的
		}(i)

		time.Sleep(time.Millisecond * 10)
	}
}

func TestGroutineB(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 发生变量共享，会输出 10 10 10 10 10 ……
		go func() {
			fmt.Println(i)
		}()

		time.Sleep(time.Millisecond * 10)
	}
}
