# Go语言示例和学习笔记

这里包含了Go语言的学习笔记和一些完整可以运行的示例，方便快速了解go语言的基本语法和特性。

# GO语言基础

* [Go语言示例和学习笔记](#go语言示例和学习笔记)
* [内容目录](#内容目录)
    * [GO背景](#go背景)
    * [GO语言基础](#go语言基础)
        * [搭建golang开发环境](#搭建golang开发环境)
        * [mac下设置GOPATH](#mac下设置gopath)
    * [变量和常量](#变量和常量)
        * [命名](#命名)
        * [变量](#变量)
        * [常量](#常量)
        * [指针](#指针)
    * [数据类型](#数据类型)
        * [整型](#整型)
        * [浮点数](#浮点数)
        * [字符串](#字符串)
        * [自定义类型](#自定义类型)
    * [集合](#集合)
        * [数组](#数组)
        * [切片 Slice](#切片-slice)
        * [Map](#map)
        * [struct 结构体](#struct-结构体)
        * [json](#json)
    * [表达式和流程控制](#表达式和流程控制)
        * [流程控制](#流程控制)
    * [错误处理](#错误处理)
    * [函数](#函数)
        * [可变参数函数](#可变参数函数)
    * [面向对象](#面向对象)
        * [方法](#方法)
        * [interface 接口类型](#interface-接口类型)
    * [包和依赖管理](#包和依赖管理)
        * [init函数](#init函数)
        * [包引入](#包引入)
    * [并发编程](#并发编程)
        * [通道（channel）](#通道channel)
    * [工具](#工具)
    * [参考文档](#参考文档)

## GO背景

应对互联网应用开发中新的挑战

1. 多核硬件架构
2. 超大规模分布式计算集群
3. Web模式导致的前所未有的开发规模和更新速度

## GO语言基础

### 搭建golang开发环境

### 配置国内代理
* go env -w GOPROXY=https://goproxy.cn,direct
* go env -u GOPROXY //取消代理

### go 常用命令
* go get -u github.com/golang/lint/golint
* go mod tidy

## 变量和常量

### 命名

* 作用于较大或生命周期较长的变量：推荐使用驼峰式命名，当名字有几个单词组成的时优先使用大小写分隔，而不是优先用下划线分隔
* 局部变量：尽量使用短小的名字，你会经常看到i之类的短名字，而不是冗长的theLoopIndex命名
* 包：建议包（package）命名用小写字母，同个目录中（不含子目录）的所有文件包名必须一致，通常为了方便包定位，建议包名和目录名一致

### 变量

创建一个变量并初始化为其零值，习惯是使用关键字var。这种用法是为了更明确地表示一个变量被设置为零值。如果变量被初始化为某个非零值，就配合结构字面量和短变量声明操作符来创建变量

```go
// var形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方
var x int = 1
var y = 1
var x int
var a, s = 100, "abc"
var {
    x, y int
    a, s = 100, "abc"
}

// 简短模式：只能在函数内部使用，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
func main() {
    x := 100
    a, s := 1, "abc"
}
```

如果初始化表达式被省略，那么将用零值初始化该变量。

1. 数值类型变量对应的零值是0，
2. 布尔类型变量对应的零值是false，
3. 字符串类型对应的零值是空字符串 "" ，
4. 接口或引用类型（包括slice、map、chan、函数）变量对应的零值是nil，
5. 数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值，

### 赋值

可赋值性的规则：类型必须完全匹配，nil可以赋值给任何指针或引用类型的变量。常量（§3.6）则有更灵活的赋值规则，因为这样可以避免不必要的显式的类型转换。

### 常量

```go
const s = "hello world"
const x, y int = 123, 0x22
const (
i, f = 1, 0.123
b = false
)
```

iota，特殊常量，可以认为是一个可以被编译器修改的常量。 iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数+1(iota 可理解为 const
语句块中的行索引)。

```go
const (
a = iota // 0
b        // 1
c        // 2
d = "ha" // 独立值，iota += 1
e        // "ha"   iota += 1
f = 100  // iota +=1
g        // 100  iota +=1
h = iota // 7,恢复计数
i        // 8
)
```

### 指针

一个指针的值是另一个变量的地址（一个指针对应变量在内存中的存储位置）。通过指针，我们可以直接读或更新对应变量的值，而不需要知道该变量的名字（如果变量有名字的话）。

```go
// &x 取地址表达式（取x变量的内存地址）
// *p 表达式对应p指针指向的变量的值
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2         // equivalent to x = 2
fmt.Println(x) // "2"
```

```go
// 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址。
p := new(int) // p, *int 类型, 指向匿名的 int 变量

// 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```

## 数据类型

在大多数高级编程语言中，数据通常被抽象为各种类型（type）和值（value）。 一个类型可以看作是值的模板。一个值可以看作是某个类型的实例。 Go支持如下内置基本类型：

* 一种内置布尔类型：bool。
* 11种内置整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。
* 两种内置浮点数类型：float32和float64。
* 两种内置复数类型：complex64和complex128。
* 一种内置字符串类型：string。

```go
*T   // 一个指针类型
[5]T // 一个元素类型为T、元素个数为5的数组类型
[]T        // 一个元素类型为T的切片类型
map[Tkey]T // 一个键值类型为Tkey、元素类型为T的映射类型

// 一个结构体类型
struct {
name string
age  int
}

// 一个函数类型
func (int) (bool, string)

// 一个接口类型
interface {
Method0(string) int
Method1() (int, bool)
}

// 几个通道类型
chan T
chan<- T
<-chan T
```

类型转换：形式为T(v)，其表示将一个值v转换为类型T，Go不支持隐式类型转换。

```go
uint(1.0)
int8(-123)
string(65)  // "A"
string('A') // "A"
string('\u68ee') // "森"
```

### 整型

11种内置整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。

Go中有两种内置类型别名（type alias）：

* byte是uint8的内置别名。 我们可以将byte和uint8看作是同一个类型。
* rune是int32的内置别名。 我们可以将rune和int32看作是同一个类型。

### 浮点数

float32 和 float64 通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大（译注：因为float32的有效bit位只有23个，
其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）：

### 字符串

* string是数据类型，不是引用或者指针类型
* string是只读的byte slice，len可以获取它包含的byte数
* 双引号用来表示字符串string，单引号表示rune类型(int32)
* Unicode是一种字符集（code point） ，UTF8是Unicode的一种存储实现，UTF8编码使用1到4个字节来表示每个Unicode point

### 自定义类型

```go
// 定义单个类型， type是关键词。
type NewTypeName SourceType

// 定义多个类型。
type (
NewTypeName1 SourceType1
NewTypeName2 SourceType2
)

// 类型等价定义，相当于类型重命名
// Celsius和Fahrenheit分别对应不同的温度单位。它们虽然有着相同的底层类型float64，但是它们是不同的数据类型
// 刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误
type Celsius float64 // 摄氏温度
type Fahrenheit float64 // 华氏温度

// 下面这些新定义的类型和它们的源类型都是组合类型。
type IntPtr *int
type Book struct{author, title string; pages int}
type Convert func (in0 int, in1 bool)(out0 int, out1 string)
type StringArray [5]string
type StringSlice []string
```

## 集合

### 数组

数组是一个由固定长度的特定类型元素组成的序列，默认情况下，数组的每个元素都被初始化为元素类型对应的零值

```go
var a[3] int
var q [3]int = [3]int{1, 2, 3} // 初始化一个数组
q := [...]int{1, 2, 3} // “...”省略号表示数组的长度是根据初始化值的个数来计算
```

### 切片 Slice

Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已 内置的append函数用于向slice追加元素

```go
// 初始化切片
s := []int{1, 2, 3 }
s1 := make([]int, 3, 5) // 类型是int，长度len=3，容量cap=5
```

### Map

Map哈希表，它是一个无序的key/value对的集合，其中所有的key都是不同的，map中所有的key都有相同的类型，所有的value也有着相同的类型

```go
/* 创建map */
countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome"}

ages := make(map[string]int) // mapping from strings to ints
ages := map[string]int{
"alice":   31,
"charlie": 34,
}
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"
```

### struct 结构体

结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。如果结构体成员名字是以大写字母开头的，那么该成员就是导出的（在其他包中可以进行读写）。

```go
type Employee struct {
ID        int
Name      string
Address   string
DoB       time.Time
}

type Point struct {
X, Y int
}

type Circle struct {
Center Point
Radius int
}

type Wheel struct {
Circle Circle
Spokes int
}

var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

### json

Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的encoding/json、encoding/xml、encoding/asn1等包提供支持.

```go
json.Marshal       // 将切片转换成json
json.MarshalIndent // 转换后的json更易读
json.Unmarshal // 解码json
```

## 表达式和流程控制

在Go中 ++ 和 -- 只能作为语句而非表达式

```go
a := 1
a ++ // 注意：不能写成 ++ a 或 -- a 必须放在右边使用
b := a++       // 此处为错误的用法，不能写在一行，要单独作为语句使用
fmt.Println(a) // 2
```

### 流程控制

Go语言中有三种基本的流程控制代码块：

* if-else条件分支代码块；
* for循环代码块；
* switch-case多条件分支代码块。

```go

for {
// 无限循环
}

for i := 0; i < 10; i++ {
fmt.Println(i)
}

// switch InitSimpleStatement; CompareOperand
switch n := rand.Intn(100); n%9 {
case 0:
fmt.Println(n, "is a multiple of 9.")
fallthrough // 跳到下个代码块,一条fallthrough语句必须为一个分支代码块中的最后一条语句
case 1, 2, 3:
fmt.Println(n, "mod 9 is 1, 2 or 3.")
break // 这里的break语句可有可无
case 4, 5, 6:
fmt.Println(n, "mod 9 is 4, 5 or 6.")
default:
fmt.Println(n, "mod 9 is 7 or 8.")
fallthrough // error: 不能出现在最后一个分支中
}
}

```

## 错误处理

* go语言没有异常机制
* error类型实现了error接口
* 可以通过 error.new来快速创建错误实例
* 及早失败，避免嵌套

* panic 用于不可恢复的错误
* panic 推出前会执行defer指定的内容

```go
type error interface (
Error() string
)

Errors.new("this is an error")

```

## 函数

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

```go
func name(parameter-list) (result-list) {
body
}
```

### 可变参数函数

“…” 是go的一种语法糖。它的两种用法：

```go
// 用法1：用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。
func test2(args ...string) { // 可变参数函数，可以接受任意个string参数
for _, v := range args{
fmt.Println(v)
}
}
// 用法2：slice可以被打散进行传递。
func main(){
var strss = []string{
"blue",
"red",
"black",
"yellow",
}

test2(strss...) // 切片被打散传入
}
```

## 面向对象

### 方法

在函数声明时，在其名字之前放上一个变量，即是一个方法。

```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// 上面的代码里那个附加的参数p，叫做方法的接收器（receiver）

p := Point{1, 2}
fmt.Println(p.Distance(q)) // "5", method call
// 这种p.Distance的表达式叫做选择器

```

当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，
这种情况下我们就需要用到指针了。对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下：

```go
func (p *Point) ScaleBy(factor float64) {
p.X *= factor
p.Y *= factor
}
```

### interface 接口类型

* 接口类型。一个接口类型定义了一个方法集，接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。
  也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。
* go允许不带任何方法的 interface ，这种类型的 interface 叫 empty interface。
    * 如果一个类型实现了一个 interface 中所有方法，我们说类型实现了该 interface， 所以所有类型都实现了 empty interface，因为任何一种类型至少实现了 0 个方法

```go
// 定义一个函数参数是 interface{} 类型，该函数可以接受任何类型作为它的参数。
func doSomething(v interface{}){
}
```

## 包和依赖管理

* 基本复用模块单元，首字母大写来表明可被外部代码访问
* 代码的package 可以和所在的目录不同
* 同一目录里的go代码的package要保持一致

### init函数

在一个代码包中，甚至一个源文件中，可以声明若干名为init的函数。 这些init函数必须不带任何输入参数和返回结果。 在程序运行时刻，在进入main入口函数之前，每个init函数在此包加载的时候将被（串行）执行并且只执行一遍。

### 包引入

句点引入：使用被句点引入的包中的导出资源时，限定标识符的前缀必须省略。一般来说，句点引入不推荐使用，因为它们会导致较低的代码可读性。

```go
package main

import (
	. "fmt"
	. "time"
)

func main() {
	Println("Current time:", Now()) // Println和Now函数调用不需要带任何前缀
}
```

匿名引入: 一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。 被匿名引入的包中的init函数将被执行并且仅执行一遍。

```go
package main

import _ "net/http/pprof"

func main() {
	... // 并没有使用匿名引入的pprof包
}
```

## GO module

Go module 构建模式是在 Go 1.11 版本正式引入的，为的是彻底解决 Go 项目复杂版本依赖的问题，在 Go 1.16 版本中，Go module 已经成为了 Go 默认的包依赖管理机制和 Go 源码构建机制。

## 并发编程

## @ Goroutines

在Go语言中，每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

goroutine很像线程，但是它占用的内存远少于线程，使用它需要的代码更少。通道（channel）是一种内置的数据结构，可以让用户在不同的goroutine之间同步发送具有类型的消息。这让编程模型更倾向于在goroutine之间发送消息，而不是让多个goroutine争夺同一个数据的使用权。

### 通道（channel）

通道（channel）是用来传递数据的一个数据结构。使用内置的make函数，我们可以创建一个channel：

```go
c := make(chan int) // 分配一个通道
// 在Go程中启动排序。当它完成后，在通道上发送信号。
go func () {
list.Sort()
c <- 1 // 发送信号，什么值无所谓。
}()
doSomethingForAWhile()
<-c // 等待排序结束，丢弃发来的值。
```

## 工具

* gofmt 保存的时候自动 格式化go代码
* goimports 保存的时候自动导入处理包 (需要先安装: go get golang.org/x/tools/cmd/goimports)
* gometalinter 保存的时候自动检查go语法

## 参考文档

* https://golang.org/pkg/ 【GO标准库中文网：英文】
* https://books.studygolang.com/The-Golang-Standard-Library-by-Example/ 【GO标准库中文网：中文】
* https://gfw.go101.org/article/101.html 【Go语言101 】
* https://github.com/xxjwxc/uber_go_guide_cn 【Uber GO语言编码规范】
* https://gorm.io/zh_CN/docs/index.html 【GORM 指南】
* github-markdown-toc 用于生成该文档的目录，#gh-md-toc /readme.md

