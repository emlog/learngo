package normal_test

import (
	"fmt"
	"strings"
	"testing"
)

type Foo struct {
	Name string
}

var bar = "hello world"

// -------------方法返回值----------------
func returnValue() string {
	return bar
}

func returnPoint() *string {
	return &bar
}

// --------------方法入参-----------------
func modifyNameByPoint(foo *Foo) {
	foo.Name = "emmmm " + foo.Name
}

func nameToUpper(foo Foo) string {
	foo.Name = strings.ToUpper(foo.Name)
	return foo.Name
}

// --------------实例方法-----------------
func (foo Foo) setName(name string) {
	foo.Name = name
}

func (foo *Foo) setNameByPoint(name string) {
	foo.Name = name
}

// 使用指针类型还是值类型，https://www.jianshu.com/p/5836fbd36dc0
func TestNormal(t *testing.T) {
	var bar = "hello biezhi"
	s1 := returnValue()
	s2 := returnPoint()

	fmt.Printf("bar: %v address: %p \n", bar, &bar) // 1
	fmt.Printf("s1 : %v address: %p \n", s1, &s1)   // 2
	fmt.Printf("s2 : %v address: %p \n", *s2, &s2)  // 3
	// output
	// » bar: hello world address: 0x115f480
	// » s1 : hello world address: 0xc00000e1e0
	// » s2 : hello world address: 0xc00000c030
	// 从这个输出中可以看到数据都是一样的，这里使用 %p 输出一个指针的值（内存地址）都不同。
	// 第一个毋庸置疑是初始的内存地址，
	// s1 是调用返回值类型的结果，
	// s2 是返回指针类型的结果。
	// 照这样看的话好像返回指针还是值类型没有区别，地址都是新的。
	//
	// 来分析一下:
	// s1 的内存地址发生变化是因为方法参数被拷贝后产生了一份新的值给 s1，此时 s1 分配了新地址。
	// s2 也拷贝了一份新值，只不过这个值是 指针类型，所以在取数据的时候用了 *。
	//
	// 既然都分配了地址，到底使用值类型还是指针类型作为返回值，我的推荐是这样的：
	//
	// 当返回类型不涉及状态变更并且是较简单的数据结构，一律返回值类型
	// 当返回类型可能遇到状态变更或者你关心它的生命周期则使用指针类型
	// 当返回的结构比较大的时候使用指针类型
}
