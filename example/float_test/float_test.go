package float_test

import (
	"fmt"
	"testing"
)

func TestFloat64(t *testing.T) {
	a := 3744
	b := 3788
	// 除法，保留两位小数的浮点型
	y := fmt.Sprintf("%.2f", float64(a)/float64(b))
	t.Log(y)
}
