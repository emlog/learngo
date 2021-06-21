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

在顶层，使用标准 var 关键字。请勿指定类型，除非它与表达式的类型不同。

```go

// bad
var _s string = F()

func F() string { return "A" }

// good
var _s = F()
// 由于 F 已经明确了返回一个字符串类型，因此我们没有必要显式指定_s 的类型
// 还是那种类型

func F() string { return "A" }
```

对于未导出的顶层常量和变量，使用 _ 作为前缀。

```go

// bad
const (
  defaultHost = "127.0.0.1"
  defaultPort = 8080
)

// good
const (
  _defaultHost = "127.0.0.1"
  _defaultPort = 8080
)
```

嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。
```go

// bad
type Client struct {
  version int
  http.Client
}

// good
type Client struct {
  http.Client

  version int
}
```

error作为函数的值返回且有多个返回值的时候，error必须是最后一个参数。

```go

// bad
func load() (error, int) {
  // normal code
}

// good
func load() (int, error) {
  // normal code
}
```

尽早进行错误处理，并尽早返回，减少嵌套。

```go

// bad
if err != nil {
  // error code
} else {
  // normal code
}

// good
if err != nil {
  // error handling
  return err
}
// normal code
```

如果需要在 if 之外使用函数调用的结果，则应采用下面的方式。

```go

// bad
if v, err := foo(); err != nil {
  // error handling
}

// good
v, err := foo()
if err != nil {
  // error handling
}
```

错误要单独判断，不与其他逻辑组合判断。

```go

// bad
v, err := foo()
if err != nil || v  == nil {
  // error handling
  return err
}

// good
v, err := foo()
if err != nil {
  // error handling
  return err
}

if v == nil {
  // error handling
  return errors.New("invalid value v")
}
```

如果返回值需要初始化，则采用下面的方式。

```go

v, err := f()
if err != nil {
    // error handling
    return // or continue.
}
// use v
```

panic 处理

* 在业务逻辑处理中禁止使用 panic。
* 在 main 包中，只有当程序完全不可运行时使用 panic，例如无法打开文件、无法连接数据库导致程序无法正常运行。
* 在 main 包中，使用 log.Fatal 来记录错误，这样就可以由 log 来结束程序，或者将 panic 抛出的异常记录到日志文件中，方便排查问题。
* 可导出的接口一定不能有 panic。
* 包内建议采用 error 而不是 panic 来传递错误。

2. 命名规范命名规范是代码规范中非常重要的一部分，一个统一的、短小的、精确的命名规范可以大大提高代码的可读性，也可以借此规避一些不必要的 Bug。

2.1 

* 包命名包名必须和目录名一致，尽量采取有意义、简短的 包名，不要和标准库冲突。
* 包名全部小写，没有大写或下划线，使用多级目录来划分层级。
* 项目名可以通过中划线来连接多个单词。
* 包名以及包所在的目录名，不要使用复数，例如，是net/utl，而不是net/urls。不要用 common、util、shared 或者 lib 这类宽泛的、无意义的包名。
* 包名要简单明了，例如 net、time、log。

2.2 函数命名

* 函数名采用驼峰式，首字母根据访问控制决定使用大写或小写，例如：MixedCaps 或者 mixedCaps。
* 代码生成工具自动生成的代码 (如 xxxx.pb.go) 和为了对相关测试用例进行分组，而采用的下划线 (如 TestMyFunction_WhatIsBeingTested) 排除此规则。

文件命名文件名要简短有意义。文件名应小写，并使用下划线分割单词。

2.4 结构体命名

采用驼峰命名方式，首字母根据访问控制决定使用大写或小写，例如 MixedCaps 或者 mixedCaps。
结构体名不应该是动词，应该是名词，比如 Node、NodeSpec。
避免使用 Data、Info 这类无意义的结构体名。
结构体的声明和初始化应采用多行，例如：

```go

// User 多行声明
type User struct {
    Name  string
    Email string
}

// 多行初始化
u := User{
    UserName: "colin",
    Email:    "colin404@foxmail.com",
}
```

2.6 变量命名

变量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。
在相对简单（对象数量少、针对性强）的环境中，可以将一些名称由完整单词简写为单个字母，例如：user 可以简写为 u；userID 可以简写 uid。
特有名词时，需要遵循以下规则：
    如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient。
    其他情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID。

若变量类型为 bool 类型，则名称应以 Has，Is，Can 或 Allow 开头，例如：

```go

var hasConflict bool
var isExist bool
var canManage bool
var allowGitHook bool
```

2.7 常量命名

常量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。如果是枚举类型的常量，需要先创建相应类型：

```go

// Code defines an error code type.
type Code int

// Internal errors.
const (
    // ErrUnknown - 0: An unknown error occurred.
    ErrUnknown Code = iota
    // ErrFatal - 1: An fatal error occurred.
    ErrFatal
)
```

2.8 Error 的命名

Error 类型应该写成 FooError 的形式。

```go

type ExitError struct {
  // ....
}
```

Error 变量写成 ErrFoo 的形式。

```go
var ErrFormat = errors.New("unknown format")
```

3. 注释规范
   
每个可导出的名字都要有注释，该注释对导出的变量、函数、结构体、接口等进行简要介绍。全部使用单行注释，禁止使用多行注释。和代码的规范一样，单行注释不要过长，禁止超过 120 字符，超过的请使用换行展示，尽量保持格式优雅。注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾，格式为 // 名称 描述. 。例如：