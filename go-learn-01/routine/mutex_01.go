package main

import (
	"fmt"
	"sync"
)

// 实验一下 资源竞争的情况 临界区
var x int64
var wg3 sync.WaitGroup
var lock sync.Mutex // 使用互斥锁
func main7() {
	wg3.Add(2)
	go add()
	go add()
	wg3.Wait()
	fmt.Println("x current value is", x)
}
func add() {
	lock.Lock()
	for i := 0; i < 1000; i++ {
		x+=1
	}
	lock.Unlock()
	wg3.Done()
}
