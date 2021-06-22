package main

// net/http提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。
// https://geektutu.com/post/gee.html

// net/http入门
import (
	"fmt"
	"log"
	"net/http"
)

// 1. 使用net/http 实现一个基本的web请求处理
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
//
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
// 获取header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}