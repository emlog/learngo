# 代码风格
   
1.1 代码格式代码都必须用 gofmt 进行格式化。运算符和操作数之间要留空格。
建议一行代码不超过 120 个字符，超过部分，请采用合适的换行方式换行。
但也有些例外场景，例如 import 行、工具自动生成的代码、带 tag 的 struct 字段。
文件长度不能超过 800 行。函数长度不能超过 80 行。

import 规范代码都必须用 goimports 进行格式化（建议将代码 Go 代码编辑器设置为：保存时运行 goimports）。
不要使用相对路径引入包，例如 import ../util/net 。
包名称与导入路径的最后一个目录名不匹配时，或者多个相同包名冲突时，则必须使用导入别名。

1.2 声明、初始化和定义当函数中需要使用到多个变量时，可以在函数开始处使用 

var 声明。在函数外部声明必须使用 var ，不要采用 := ，容易踩到变量的作用域的问题。

```go
var ( 
    Width int 
    Height int
)
```

在初始化结构引用时，请使用 &T{}代替 new(T)，以使其与结构体初始化一致。
```go

// bad
sval := T{Name: "foo"}

// inconsistent
sptr := new(T)
sptr.Name = "bar"

// good
sval := T{Name: "foo"}

sptr := &T{Name: "bar"}
```

struct 声明和初始化格式采用多行，定义如下。
```go
type User struct{
    Username  string
    Email     string
}

user := User{
    Username: "colin",
    Email: "colin404@foxmail.com",
}
```

尽可能指定容器容量，以便为容器预先分配内存，例如：

```go
v := make(map[int]string, 4)
v := make([]string, 0, 4)
```

