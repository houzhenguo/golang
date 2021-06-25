package main

import "fmt"

func main() {
	var s1 = a1{}
	var s2 = b1{}
	testFuc(s1)
	testFuc(s2)
	var a interface{} // 定义空接口，必须要用上，否则就提示你转 _
	b := int32(2)
	a = b

	// 可以判断接口的值
	_, ok1 := a.(string)
	if ok1 {
		fmt.Println("string ", a)
	}
	_, ok2 := a.(int32) // 断言
	if ok2 {
		fmt.Println("int32 ", a)
	}
	fmt.Println(a)
	// 空interface 类似 Java中都object

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "houzhenguo"
	studentInfo["age"] = 12
	studentInfo["score"] = 99.9
	fmt.Println(studentInfo)
}

type a1 struct {

}
type b1 struct {

}
type bark interface {
	sayHello()
}

func (a a1) sayHello()  {
	fmt.Println("a1 say")
}
func (b b1) sayHello()  {
	fmt.Println("b1 say")
}
func testFuc(s say)  {
	s.sayHello()
}
