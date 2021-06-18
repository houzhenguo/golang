package main

import (
	"fmt"
	"os"
)
// ./args a b c
// build 一下启动
func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("arsgs[%d]=%v\n", index, arg)
		}
	}
}
