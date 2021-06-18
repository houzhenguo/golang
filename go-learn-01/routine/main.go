package main

import (
	"fmt"
	"sync"
)

// golang 协程 单个routine 占用 2k,相比线程 2M 小很多
var wg sync.WaitGroup // 使用wg 等待
var wg1 sync.WaitGroup

func main1() {
	fmt.Println("hello goroutine")
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	wg.Wait()

	// 匿名函数
	wg1.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			wg1.Done()
			fmt.Println("匿名", i) // i 如果不处理会重复，闭包
		}(i)
	}
	wg1.Wait()
	fmt.Println("hello main")
}
func hello(i int) {
	fmt.Println("hello one", i)
	wg.Done()
}
