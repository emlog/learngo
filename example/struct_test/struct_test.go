/*
结构体（struct）的使用
*/
package struct_test

import (
	"testing"
)

// 初始化切片
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

	t.Logf("user: %s, email: %s", user.Username, a.Email)
}
