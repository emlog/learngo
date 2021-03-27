package groutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i) // 输出顺序可能不是顺序的1234567…… 是因为协程的调用并不是顺序的，存在并发
		}(i)
	}
	time.Sleep(time.Second * 2)
	t.Log("done")
}

func TestGroutineB(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 发生变量内容共享，会输出 10 10 10 10 10 ……
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
	t.Log("done")
}

// 非线程安全的计数方式，计数可能不准
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Logf("counter is : %d", counter) // 计数结果可能小于5000
}

// 使用sync.wait 来等待所有的groutine执行完成。不使用sleep的方式
// 使用sync.mutex 来锁定共享内存，确保计数准确。
func TestCounterThreadeSafe(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}

	wg.Wait()
	t.Logf("counter is : %d", counter) // 计数结果可能小于5000
}
