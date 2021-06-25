package main

import (
	"fmt"
	"strings"
)

func main() {
	// ReaderAt 练习 io.ReaderAt
	reader := strings.NewReader("问问振国爱学习")
	p := make([]byte, 19)
	n, err := reader.ReadAt (p, 0)
	fmt.Printf("n is", n)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d \n", p, n)

	//p1 := make([]byte, 2)
	//bytes := []byte("问问振国爱学习")
	//n := copy(p1, bytes[2:])
	//fmt.Printf("n is %d, and %s, and length is %d and bytes lens is %d \n", n, p1, len(p1), len(bytes))
}
