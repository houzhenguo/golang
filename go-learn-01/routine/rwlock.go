package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁 读读不加锁 ，读写，写写 加锁
var (
	x5     int64
	wg5    sync.WaitGroup
	rwlock sync.RWMutex // 使用读写锁 耗时 113.979742ms
	lock5  sync.Mutex // 使用 互斥锁 耗时 1.390447392s  差了10倍
)

func main8() {
	fmt.Println("read write lock ")
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg5.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg5.Add(1)
		go read()
	}
	wg5.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}
func write() {
	//rwlock.Lock() // rwlock 使用 lock方法
	lock5.Lock()
	x5 += 1
	time.Sleep(10 * time.Millisecond)
	//rwlock.Unlock()
	lock5.Unlock()
	wg5.Done()
}
func read() {
	//rwlock.RLock()   注意加锁/解锁的方法
	lock5.Lock()
	time.Sleep(time.Millisecond)
	//rwlock.RUnlock()
	lock5.Unlock()
	wg5.Done()
	//fmt.Printf()
}
