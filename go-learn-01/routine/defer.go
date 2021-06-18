package main

import "fmt"

// test defer
func main9() {
	defer  work() // defer 类似Java中的finally
	fmt.Println("hello ")
}

func work(){
	fmt.Println("work")
}