## golang学习笔记
这里包含了GO语言的学习笔记和一些可以运行的示例

### mac下设置GOPATH

vi ~/.zshrc 添加以下代码 (永久设置环境变量)

```shell script
export GOROOT=/usr/local/go
export GOPATH=/Users/xudawei/gowork
export GOBIN=
export PATH=$PATH:${GOPATH//://bin:}/bin
export GO111MODULE=auto
```

* GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
* GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
* GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
* 当前目录在GOPATH/src之外且该目录包含go.mod文件
* 当前文件在包含go.mod文件的目录下面。


### 命名
* 作用于较大或生命周期较长的变量：推荐使用驼峰式命名，当名字有几个单词组成的时优先使用大小写分隔，而不是优先用下划线分隔
* 局部变量：尽量使用短小的名字，你会经常看到i之类的短名字，而不是冗长的theLoopIndex命名
* 包：包的名字应只用小写（不要用下划线式，也不要用驼峰式）

### 声明
Go语言主要有四种类型的声明语句：

* var：变量
* const：常量
* type：类型
* func：函数实体对象

### 变量

```go
//var形式的声明语句往往是用于需要显式指定变量类型地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方
var x int
var y = false
var x, y int 
var a, s = 100, "abc"
var {
    x,y int
    a,s = 100, "abc"
}
//简短模式：只能在函数内部使用
//:=，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
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

### 常量
运行时恒定不改变的值

```go
const s = "hello world"
const x, y int = 123, 0x22
const (
    i, f = 1, 0.123
    b = false
)
```
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数+1(iota 可理解为 const 语句块中的行索引)。
```go
const (
        a = iota   //0
        b          //1
        c          //2
        d = "ha"   //独立值，iota += 1
        e          //"ha"   iota += 1
        f = 100    //iota +=1
        g          //100  iota +=1
        h = iota   //7,恢复计数
        i          //8
)
fmt.Println(a,b,c,d,e,f,g,h,i)
//运行结果：0 1 2 ha ha 100 100 7 8
```

### channel  chan 管道
ch := make(chan int) //创建了一个没有缓冲区，只能读写int数据类型的channel

### 指针
一个指针的值是另一个变量的地址（一个指针对应变量在内存中的存储位置）。通过指针，我们可以直接读或更新对应变量的值，而不需要知道该变量的名字（如果变量有名字的话）。
```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```
&x 取地址表达式（取x变量的内存地址）
*p 表达式对应p指针指向的变量的值

表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址。
p := new(int)   // p, *int 类型, 指向匿名的 int 变量

每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"


### type，自定义类型
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
我们在这个包声明了两种类型：Celsius和Fahrenheit分别对应不同的温度单位。它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，
因此它们不可以被相互比较或混在一个表达式运算。刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误。

//类型等价定义，相当于类型重命名
type name string
name类型与string等价

自定义一个结构体类型：
```go
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
}
```

### 整型
有符号整形：int8、int16、int32、int64
无符号整形：uint8、uint16、uint32、uint64
var u uint8 = 255
var i int8 = 127

### 浮点数
float32  和 float64 
通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大（译注：因为float32的有效bit位只有23个，
其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）：

### 字符串
子字符串操作s[i:j] 基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串
```go
s := "hello, world"
fmt.Println(len(s))     // "12"
fmt.Println(s[0]) // "104" ('h' 索引操作s[i]返回第i个字节的字节值)
fmt.Println(s[0:5]) // "hello"
fmt.Println(s[0:1]) // "h"
```

UTF8编码使用1到4个字节来表示每个Unicode码点

### 数组
数组是一个由固定长度的特定类型元素组成的序列，默认情况下，数组的每个元素都被初始化为元素类型对应的零值
var a[3] int
var q [3]int = [3]int{1, 2, 3} 初始化一个数组
q := [...]int{1, 2, 3}   //在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算


### 切片 Slice
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已
内置的append函数用于向slice追加元素
```go
//定义切片变量
var identifier []type
//初始化切片
s :=[] int {1,2,3 } 
```

### Map
哈希表，它是一个无序的key/value对的集合，其中所有的key都是不同的，map中所有的key都有相同的类型，所有的value也有着相同的类型
```go
/* 创建map */
countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
```

```go
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
json.Marshal //将切片转换成json
json.MarshalIndent //转换后的json更易读
json.Unmarshal // 解码json
```

### 函数

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
```go
func name(parameter-list) (result-list) {
    body
}
```

### 方法
在函数声明时，在其名字之前放上一个变量，即是一个方法。
```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
//上面的代码里那个附加的参数p，叫做方法的接收器（receiver）

p := Point{1, 2}
fmt.Println(p.Distance(q))  // "5", method call
//这种p.Distance的表达式叫做选择器

```

当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，
这种情况下我们就需要用到指针了。对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法，如下：
```go
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
```

### 接口
接口类型。接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。
也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

### Goroutines
在Go语言中，每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

### 通道（channel）
通道（channel）是用来传递数据的一个数据结构。使用内置的make函数，我们可以创建一个channel：
```go
ch := make(chan int) // ch has type 'chan int'
```

### 工具链

测试框架：单元测试、性能测试、代码覆盖率、PPROF

### 相关文档

* https://golang.org/pkg/ 【GO标准库中文网：英文】
* https://books.studygolang.com/The-Golang-Standard-Library-by-Example/ 【GO标准库中文网：中文】
* https://gorm.io/zh_CN/docs/index.html 【GORM 指南】












