// 并发模式：只运行一次，（单例模式，在多线程环境下某一块代码只执行一次）
package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	data string
}

var singleInstance *Singleton
var once sync.Once

// 单例的实现
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// 多线程环境下确保只执行一次
func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//打印对象地址 十六进制
			fmt.Printf("%x\n", unsafe.Pointer(obj)) // %x:base 16, with lower-case letters for a-f
			wg.Done()
		}()
	}
	wg.Wait()
}
