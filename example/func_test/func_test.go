/*
函数的使用
*/
package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 函数作为参数，计算函数的执行时间
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Snow() (a int, err error) {
	a = 111
	err = nil

	return
}

func TestVarParam2(t *testing.T) {
	b2, err := Snow()

	fmt.Printf("b2: %d, err: %v", b2, err)
}

func Foo() {
	if e := recover(); e != nil {
		fmt.Println("recovered from ", e)
	}

	panic("panic err!!")
}
