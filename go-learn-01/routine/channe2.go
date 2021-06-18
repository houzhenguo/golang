package main

import (
	"fmt"
	"time"
)

func main3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			} else {
				//ch2 <- i
				fmt.Println(i)
			}
		}
		close(ch2)
	}()
	time.Sleep(1 * time.Second)
	//for i:= range ch2 {
	//	fmt.Println("i is", i)
	//}

	// 还有单向通道
}
