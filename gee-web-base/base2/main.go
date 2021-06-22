package main

import (
	"fmt"
	"log"
	"net/http"
)

// 重写 httpHandler

type Engine struct {

}
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
// 实现一下 Handler#ServeHTTP 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
