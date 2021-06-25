package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	// 这个不能go function 里面，add 放在外面
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	fmt.Println("main finish")
}
func hello(group *sync.WaitGroup)  {

	time.Sleep(2 * time.Second)
	fmt.Println("hello routine")
	defer group.Done()
}
