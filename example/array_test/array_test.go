// 数组（array）的使用
package array_test

import (
	"fmt"
	"testing"
)

// 数组初始化
func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4} // "..." 表示数组的长度是根据初始化值的个数来计算
	arr3 := [2]string{"snow", "sun"}
	t.Log(arr[1], arr[2])
	fmt.Println(arr1, arr2, arr3)

	// 二维数组
	var arr4 = [2][4]int{{1, 2, 3}, {4, 5, 6}}
	t.Log(arr4[0][1])
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{3, 4, 6, 7, 8}

	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// range
	for k, v := range arr3 {
		t.Log(k, v)
	}
}

// 截取数组中的第一段
func TestArraySection(t *testing.T) {
	arr3 := [...]int{3, 4, 6, 7, 8}
	arrSec1 := arr3[3:]
	arrSec2 := arr3[:5]
	arrSec3 := arr3[1:4]
	t.Log(arrSec1, arrSec2, arrSec3)
}
