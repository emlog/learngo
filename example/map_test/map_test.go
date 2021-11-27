// map哈希表的使用
package map_test

import (
	"testing"
)

// 初始化map
func TestInitMap(t *testing.T) {
	// map [key]value
	// value的类型没有限制，但是key 的类型必须支持“==”和“!=”两种比较操作符，函数
	// 函数类型、map 类型自身，以及切片类型是不能作为 map 的 key 类型的
	m1 := map[int]int{1: 2, 2: 3, 3: 4}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[1] = 1
	t.Logf("len m2=%d", len(m2))

	// 使用make来初始化map，指定初始容量为8
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

	// map1 is nil
	var map1 map[string]int
	// map2 is not nil
	map2 := map[string]int{}
	if map2 != nil {
		t.Log("map1 not nil")
	}
	// 初值为零值 nil 的切片类型变量，可以借助内置的 append 的函数进行操作，这种在 Go 语言中被称为“零值可用”
	// 但 map 类型，因为它内部实现的复杂性，无法“零值可用”
	// panic: assignment to entry in nil map
	map1["key"] = 1
	// no panic
	map2["key"] = 1

}

// 区分key不存在和0值
func TestAccessNoexistKey(t *testing.T) {
	m1 := map[int]int{}

	t.Log(m1[2]) // key 2 不存在返回0

	m1[3] = 0
	t.Log(m1[3]) // key 3 存在但是是0值，也返回0

	// 区分不存和0值
	if v, ok := m1[3]; ok {
		t.Logf("key 3 's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestDelMapKey(t *testing.T) {
	m := map[string]int{"key1": 1, "key2": 3}
	// 删除"key2",delete 函数的执行也不会失败，更不会抛出运行时的异常
	delete(m, "key2")
	t.Logf("len of m:%d", len(m))
}

// 遍历map
// map 是无序的,每次打印出来的 map 都会不一样,它不能通过 index 获取,而必须通过 key 获取
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3, 3: 4, 4: 5}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

// 函索作为map的值
func TestMapWtihFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// 用map实现集合set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3

	// 判断集合中是否存在某值
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	// 集合长度
	t.Log(len(mySet))

	// 删除一个key
	delete(mySet, 1)

}
