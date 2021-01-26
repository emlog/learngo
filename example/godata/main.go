/*
代码示例：数据类型
*/
package main

import "fmt"

// 定义数组
var colorsList [3]string
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// 忽略长度的定义
var balance2 = [...]string{"aaa", "bbb", "ccc"}

// 定义切片
var slice1 []string

// Books 定义结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 数组
	animals := [3]string{"cat", "dog", "monkey"}
	for i := 0; i <= 2; i++ {
		fmt.Printf("animal is %s \n", animals[i])
	}

	// 切片
	slice2 := []string{"red", "blue", "green"}
	fmt.Printf("item0: %s , slice2 is len: %d \n", slice2[0], len(slice2)) // len函数获取切片的长度
	for _, value := range slice2 {                                         // 遍历切片
		fmt.Printf("%s \n", value)
	}

	// map
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["China"] = "北京"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	/*
		countryCapitalMap := map[string]string {
			"China" : "beijing",
			"Italy" : "luoma",
			"japan" :"dongjing",
		}
	*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

}
