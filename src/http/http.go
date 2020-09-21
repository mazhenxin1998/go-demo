package main

import (
	"fmt"
	"net/http"
)

func main() {


}

func init() {
	http.HandleFunc("/get/go/say",request)
	// 设置监听端口.
	http.ListenAndServe("127.0.0.1:8080",nil)

}

/**
简单的HTTP请求,对比Java中的servlet.
 */
func request(w http.ResponseWriter, r *http.Request) {

	// 循环10000次进行输出
	for i := 0; i < 10000; i++ {

		fmt.Println("1")
	}

	str := "你好,GoLang"
	// 发送的时候需要将字符串转换成字节数组进行传输.
	w.Write([]byte(str))

}
