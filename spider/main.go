package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 爬邮箱
func GetEmail() {
	// 1.get data
	resp, err := http.Get("https://booltool-pre.boolv.tech/home")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 解析页面内容
	body, err := ioutil.ReadAll(resp.Body)
	stringBody := string(body)
	fmt.Println(stringBody)
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	GetEmail()
}
