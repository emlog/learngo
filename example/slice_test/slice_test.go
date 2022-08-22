// 切片（slice）的使用
package slice_test

import (
	"sort"
	"testing"
)

// 初始化切片
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) // 打印切片长度len、容量cap

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 3, 5}
	t.Log(len(s1), cap(s1))

	t.Logf("%v", s1) // [ 1 3 5 ]

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	s2 = append(s2, 1)
	t.Log(s2[1], s2[3])

	// nil切片在运行时表示的三个字段值都是0；而空切片的len、cap值为0，但data值不为零。
	// 空切片到底分没分配底层数组？ 答案是肯定的：没有分配！
	var sl1 = []int{} // sl1是空切片
	var sl2 []int     // sl2是nil切片

	for x, y := range sl1 {
		t.Logf("empty slice: x:%d, y:%d", x, y)
	}

	for x, y := range sl2 {
		t.Logf("nil slice: x:%d, y:%d", x, y)
	}

}

// 遍历切片
func TestLoopOverSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		t.Log(v)
	}
}

// 二维切片的初始化，排序
func Test2DSlice(t *testing.T) {
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sort.Slice(twoD, func(i, j int) bool {
		if len(twoD[i]) <= 3 || len(twoD[j]) <= 3 {
			return false
		}
		return twoD[i][3] > twoD[j][3]
	})
	t.Log(twoD)
}

// 2个切片的 append 追加
func TestSliceAppend(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5}
	var b = []int{6, 7, 8, 9, 10}
	s = append(s, b...)
	t.Log(s)
}

// 切片长度自动扩容，前一次的2倍
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片共享内容
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unKonw"
	t.Log(Q2)
	t.Log(summer)
	t.Log(year)
}

// 切片比较
func TestSliceComparing(t *testing.T) {
	a := []int{1, 3, 5}

	// b := []int{1, 3, 5}
	// slice can only be compared to nil
	// if a == b {
	// 	t.Log("equal")
	// }

	if a == nil {
		t.Log("equal")
	}
}
