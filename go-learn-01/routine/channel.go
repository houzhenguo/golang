package main

import (
	"fmt"
	"time"
)

// channel 在并发中交换数据才有意义
func main2() {
	var c1 chan int
	c1 = make(chan int, 1) // 构建一个放int 的chan
	c1 <- 10
	go recv(c1)
	//a1 := <-c1
	//fmt.Println("a1 接收到到值", a1)
	// close(c1)
	// c1 <- 30 关闭之后
	time.Sleep(10 * time.Second)
}
func recv(c chan int) {
	v := <-c
	fmt.Println("接收v", v)
}
