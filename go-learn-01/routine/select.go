package main

import "fmt"

// select 语法
// select 可以处理一个或者多个channel的发送/接收操作
// 对于 多个case select 会随机挑一个
func main6() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("x is", x)
		case ch <- i:
		}
	}
}
