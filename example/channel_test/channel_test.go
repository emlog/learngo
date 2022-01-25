package channel_test

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// var ch1 chan int
	// ch1 := make(chan int, 5) 带缓冲区的channel
	ch1 := make(chan int)
	go func() {
		// 等待channel有数据
		time.Sleep(time.Second * 2) // Duration的单位为 nanosecond
		ch1 <- 100
	}()

	a := <-ch1
	t.Logf("a:%d", a)
}
