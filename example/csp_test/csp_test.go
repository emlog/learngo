/*
CSP 并发机制的使用。
Go 语言实现了基于 CSP（Communicating Sequential Processes）理论的并发方案。
Go 语言的 CSP 模型的实现包含两个主要组成部分：一个是 Goroutine，它是 Go 应用并发设计的基本构建与执行单元；另一个就是 channel，它在并发模型中扮演着重要的角色。
channel 既可以用来实现 Goroutine 间的通信，还可以实现 Goroutine 间的同步。在 Go 并发设计时灵活运用 channel，才能说真正掌握了 Go 并发设计的真谛。
*/
package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 5000)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// channel
func AsyncService() chan string {

	// 1、阻塞，初始化一个channel，类型是字符串
	// retCh := make(chan string)

	// 2、非阻塞的buffer通道，初始化一个带buffer的channel, 信息放入通道的buffer后 该groutine会推出，而不是一直等待消息被接收，是更高效的写法。
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret // 传递一个字符串信息到channel
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 接收并打印一个字符串的channel信息
	time.Sleep(time.Second * 1)
}

// select 多路选择的测试
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Second * 1): // 超过1s后报超时错误
		t.Error("time out ")
	}
}

// 通道基础测试
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

func TestA(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick // 接收一个节拍器的事件。
	}
}
