package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)
// 漏斗限流 -> 不能应对突发流量处理 -> 按照固定的流量处理
// #atomicLimiter
// https://www.liwenzhou.com/posts/Go/ratelimit/
func main() {
	rl := ratelimit.New(100)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
	}
}
