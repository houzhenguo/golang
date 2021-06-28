package main

import (
	"fmt"
	"sync"
)

// mutex.lock + sync.WaitGroup 配合使用
func main2() {
	count := 0
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			count += 1
			wg.Done()
			lock.Unlock()
		}()
	}
	//time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println("count value is", count)

}
