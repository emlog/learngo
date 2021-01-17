package example

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name string `json:"name"` //大写开头可导出可以被转为json，
	Age  int8   `json:"age"`
	sex  string `json:"sex"` //小写字母开头，不可以被转为json
}

func main() {

	//结构体转json数据
	m := userInfo{
		"snow",
		10,
		"boy",
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b)) //b是[]byte类型，转化成string类型便于查看

	//json转结构体
	//json字符中的"引号，需用\进行转义，否则编译出错
	data := "{\"name\":\"张三\",\"Age\":18,\"sex\":\"男\"}"
	str := []byte(data)

	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	stu := userInfo{}

	err = json.Unmarshal(str, &stu)
	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(stu.Name)

}
