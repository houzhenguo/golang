package main

import "fmt"

// 在写排序的时候遇到一个函数接口练习证明一下
func main1() {
	//f := func(a, b int) int {
	//	return a + b
	//}
	//FuncTest(f).test() // 只有是个方法的时候将 这个 f 相当于构造 放入到里面

	// 测试
	// 检测 是否实现了 say interface ,如果没有完全实现，
	// 则会在编译期就报错
	// var _ say = bird{} // 我先注释掉
	var _ say = &cat{}

	// method中是结构体，golang语法糖会字段转成 *cat
	// 如果指定都是 *cat,则不能传递 struct

}

type FuncTest func(a, b int) int

func (f FuncTest) test() {

}

type cat struct {

}
type dog struct {

}
type bird struct {

}
type say interface {
	sayHello()
}

func (c cat) sayHello()  {
	fmt.Printf("cat say")
}
func (d dog) sayHello()  {
	fmt.Printf("dog say")
}