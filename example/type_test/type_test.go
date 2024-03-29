// 数据类型的使用和定义
package type_test

import (
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") // 初始化零值是“”
	t.Log(len(s))

}

// 类型断言测试
func TestAssertion(t *testing.T) {
	var x interface{}
	x = "10"
	value, ok := x.(int)
	t.Log(value, ",", ok)
}
