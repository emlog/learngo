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
	arr := [...]int{99: 1} // 定义有100个元素的数组arr，除最后一个是1外，其他都是0，没有指定值得索引位置被赋予类型零值

	// range
	for k, v := range arr {
		t.Log(k, v)
	}

	arr1 := [...]int{2, 3, 4, 5, 6}

	// range
	for k, v := range arr1 {
		t.Log(k, v)
	}
}

// 截取数组中的第一段
func TestArraySection(t *testing.T) {
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arrSec1 := arr3[1:4] // 从下标1开始，到下标4结束（包括下标1，不包括下标4）
	arrSec2 := arr3[3:]
	arrSec3 := arr3[:5]
	t.Log(arrSec1, arrSec2, arrSec3)
}
