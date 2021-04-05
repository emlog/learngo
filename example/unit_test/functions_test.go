package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 不使用断言 assert
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d , the expected is %d", inputs[i], expected[i])
		}
	}
}

// 使用相等断言 assert
func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}

// Error
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

// Fatal
func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

/*
显示单元测试代码覆盖度
go test -v -cover
*/
