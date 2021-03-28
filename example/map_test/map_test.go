// map哈希表的使用
package map_test

import "testing"

// 初始化map
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3, 3: 4}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[1] = 1
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

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

// 遍历map
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3, 3: 4}
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
