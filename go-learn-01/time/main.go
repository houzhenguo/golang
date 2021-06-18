package main

import (
	"fmt"
	"time"
)

var counter int64
func main() {
	//ticker := time.Tick(time.Second * 2) // ticker 每隔两秒钟执行一次
	//for i := range ticker {
	//	counter +=1
	//	fmt.Println("hello", counter)
	//	fmt.Println("i is", i)
	//}
	formatDemo()
}

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
