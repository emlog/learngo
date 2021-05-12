/*
字符串的使用
*/
package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)    // 初始化为默认零值“”
	s = "hello" // 在Go中，双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型(int32)
	t.Log(len(s))
	// s[1] = '3' //string是不可变的byte slice
	// s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) // 是byte数

	c := []rune(s)
	t.Log(len(c))
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// %c	the character represented by the corresponding Unicode code point
		// %x	base 16, with lower-case letters for a-f
		t.Logf("%[1]c %[1]x", c) // [1]表示第一个参数c，分别按照%c 和 %x格式化.
	}
}

// 字符串分割 split 和连接 join
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
	t.Log(strings.Join(parts, "+++"))
}

// 字符串和int的转换
func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	} else {
		t.Log("can not conv")
	}
}

func TestInt2Str(t *testing.T) {
	var uid int
	uid = 123456

	// 数值类型和 string 类型之间的相互转换可能造成值部分丢失
	// uidstring := string(uid)

	// string 和数字之间的转换可使用标准库 strconv
	uidstring := strconv.Itoa(uid)

	t.Logf("uid: %s", uidstring)

}

func TestBenchmarkOk(t *testing.T) {
	var str = "1"
	var number int = 2
	fmt.Println("原字符串：", str)
	fmt.Println("原数字：", number)
	// 字符串转数字
	num, _ := strconv.Atoi(str)
	fmt.Println("字符串转数字：", num)
	// 数字转字符串
	str1 := strconv.Itoa(num)
	fmt.Println("数字转字符串：", str1)
	// 数字转数字
	var numTrans = int64(number)
	fmt.Println("数字转数字：", numTrans)
}
