/*
结构体（struct）的使用
*/
package struct_test

import (
	"testing"
)

// 初始化一个结构体
func TestStructInit(t *testing.T) {
	// 声明结构体类型
	type User struct {
		Username string
		Email    string
	}

	// 初始化结构引用
	a := &User{}
	a.Email = "emlog@qq.com"

	// 初始化结构体
	user := User{
		Username: "colin",
		Email:    "colin404@foxmail.com",
	}

	//	when printing structs, the plus flag (%+v) adds field names
	t.Logf("%+v", user)

}
