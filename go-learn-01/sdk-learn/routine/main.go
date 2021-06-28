package main

import (
	"fmt"
	"sync"
	"time"
)

func main3() {
	//var wg = sync.WaitGroup{}
	//// 这个不能go function 里面，add 放在外面
	//wg.Add(1)
	//go hello(&wg)
	//wg.Wait()
	//fmt.Println("main finish")

	testChan()
}
func hello(group *sync.WaitGroup) {

	time.Sleep(2 * time.Second)
	fmt.Println("hello routine")
	defer group.Done()
}

func testChan() {
	var ch1 chan int
	ch1 = make(chan int, 10) // channel 必须make之后才能使用
	write(ch1)
	go read(ch1) // read and write 不能在同一个线程中
	defer close(ch1)
	time.Sleep(3 * time.Second)
	fmt.Println("ch1", ch1)
}

func write(ch1 chan int) {
	ch1 <- 10
	ch1 <- 20
}
func read(ch1 chan int) {
	for v := range ch1 {
		fmt.Println(v)
	}
}

//func singleChannel(c chan <- int) {
//	a := <-c // 单向通道，这里会报错
//	fmt.Println(a)
//}
