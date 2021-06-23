package main

import "fmt"

func main() {
	testGet()
}
type String string

func (d String) Len() int {
	return len(d)
}
func testGet()  {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		fmt.Println("cache hit key1=1234 failed")
	}
	fmt.Println("hello")
	if _, ok := lru.Get("key2"); ok {
		fmt.Println("cache miss key2 failed")
	}
}
