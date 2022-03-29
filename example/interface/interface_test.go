// 接口的使用
package interface_test

import (
	"fmt"
	"testing"
)

// 定义一个接口 Programmer
type Programmer interface {
	WriteHelloWorld() string
}

// 定义接口实现
type GoProgrammer struct {
}

// 实现接口 Programmer
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

type E interface {
	M1()
	M2()
}

// 像这种在一个接口类型（I）定义中，嵌入另外一个接口类型（E）的方式，就是我们说的接口类型的类型嵌入。
type I interface {
	E
	M3()
}

type T1 int
type t2 struct {
	n int
	m int
}

type J interface {
	M1()
}

type S1 struct {
	T1
	*t2
	J
	a int
	b string
}

// getType 类型断言 返回接口类型的实际类型
func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}

func TestInterface(t *testing.T) {
	getType(1)
	getType("hello")
	getType(1.1)
}
