/*
代码示例：http调用
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	resp, err := http.Get("https://emlog.net")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
	fmt.Printf("status code: %d", resp.StatusCode)
}

func main() {
	httpGet()
}
