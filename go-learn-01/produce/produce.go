package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用go 实现简易的生产者消费者模式
func main() {
	ch1 := make(chan string, 10)
	go produce("dog", ch1)
	go produce("cat", ch1)
	go consume(ch1)
	time.Sleep(30 * time.Second)
}
// 生产者
func produce(prefix string, c chan<- string) {
	for  {
		s := fmt.Sprintf("%s:%v", prefix, rand.Int31n(100))
		c<-s
		time.Sleep(time.Second)
	}
}
// 消费者
func consume(c <-chan string) {
	for {
		v:= <-c
		fmt.Println("v is", v)
	}
}
