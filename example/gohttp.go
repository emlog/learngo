package example

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	httpGet()
	httpDo()
}

func httpGet() {
	resp, err := http.Get("https://www.emlog.net/license")
	if err != nil {
		fmt.Println("error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://www.emlog.net/license", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println("error")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(string(body))
}
