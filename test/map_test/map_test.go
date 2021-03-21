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
