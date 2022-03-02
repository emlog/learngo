// 常量的使用
package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota // 位运算
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 // 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
