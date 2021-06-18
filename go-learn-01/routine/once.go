package main

import (
	"fmt"
	"sync"
	"time"
)

// singleton
var instance *singleton
var once sync.Once
func main10() {
	a:=3
	b:= &a
	fmt.Println("a address", &a)
	fmt.Println("b value", b) // b == &a 他们两个的值是相同的
	fmt.Println("b address", &b) // 存储地址也需要内存


	go func() {
		v:= GetInstance()
		fmt.Println("v address",&v)
	}()
	go func() {
		v1:= GetInstance()
		fmt.Println("v1 address",&v1)
	}()
	time.Sleep(3 * time.Second)
}
func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("do once")
		instance = &singleton{
			name: "zz",
		}
	})
	fmt.Println("instance address",&instance)
	return instance
}
type singleton struct {
	name string
}