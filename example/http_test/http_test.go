/*
http 网络请求的使用
*/
package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
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

func httpPost() {
	urlValues := url.Values{}
	urlValues.Add("name", "snow")
	urlValues.Add("age", "22")
	resp, _ := http.PostForm("http://httpbin.org/post", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestHttp(t *testing.T) {
	httpGet()
	httpPost()
}
