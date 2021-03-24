package array_test

import "testing"

// 数组初始化 init array
func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4} // "..."表示数组的长度是根据初始化值的个数来计算
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

// 数组遍历 array travel
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
	arr_sec1 := arr3[3:]
	arr_sec2 := arr3[:5]
	arr_sec3 := arr3[1:4]
	t.Log(arr_sec1, arr_sec2, arr_sec3)
}
