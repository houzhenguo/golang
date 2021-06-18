package main

import (
	"fmt"
	"go-learn-01/gint/routers"
)

func main2()  {
	fmt.Println("hello")
	//
	r:=routers.SetupRouter()
	if err := r.Run(":9090"); err != nil {
		fmt.Println("start up failed")
	}
}
