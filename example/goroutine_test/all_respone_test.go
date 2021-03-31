// 并发模式：所有任务都完成
package goroutine_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask2(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

// 等到接收到所有的消息才返回
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, 10)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask2(i)
			ch <- ret
		}(i)
	}
	finalret := ""
	for i := 0; i < numOfRunner; i++ {
		finalret += <-ch + "\n"
	}
	return finalret
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 返回当前运行时的goroutine数量
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
