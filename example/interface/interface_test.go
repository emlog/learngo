package interface_test

import "testing"

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
