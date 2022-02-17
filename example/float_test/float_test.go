package float_test

import (
	"fmt"
	"testing"
)

func TestFloat64(t *testing.T) {
	var a float64 = 3744
	var b float64 = 3788
	// 除法，保留两位小数的浮点型
	y := fmt.Sprintf("%.2f", a/b)
	t.Log(y)
}
