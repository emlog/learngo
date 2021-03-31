// 并发模式：任意任务完成就可以
package goroutine_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

// 接收到第一个消息就返回
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, 10)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 返回当前运行时的goroutine数量
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
