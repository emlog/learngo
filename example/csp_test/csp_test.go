/*
CSP 并发机制的使用。
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

	//1、阻塞，初始化一个channel，类型是字符串
	//retCh := make(chan string)

	//2、非阻塞的buffer通道，初始化一个带buffer的channel, 信息放入通道的buffer后 该groutine会推出，而不是一直等待消息被接收，是更高效的写法。
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //传递一个字符串信息到channel
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //接收并打印一个字符串的channel信息
	time.Sleep(time.Second * 1)
}

// select 多路选择的测试
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Second * 1): //超过1s后报超时错误
		t.Error("time out ")
	}
}
