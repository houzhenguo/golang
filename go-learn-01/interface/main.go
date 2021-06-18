package main

import "fmt"

func main() {
	fmt.Println("hello interface")
	d1 := &dog{   // 对于指针类型的必须加 &
		name: "d1",
	}
	var c say   // 如果想要 say 接收，必须实现say的所有接口
	c = d1
	fmt.Println(c)
	d1.say(d1.name)
	d1.eat()

	c1 := cat{
		name: "c1",
	}
	c1.say(c1.name)
	// 测试
	null_inter("lisi")
	null_inter(22)

	// 测试接口switch case
	null_switch_interface("224")
	null_switch_interface(44)
	null_switch_interface(int32(33))
	null_switch_interface(d1) // 类似 Java中的object
}
type dog struct {
	name string
}

type cat struct {
	name string
}
type say interface {
	say(name string)
	eat()
}

func (d dog) say(name string)  {
	fmt.Println("I'm a dog ..name is ", name)
}
func (d dog)eat()  {
	fmt.Println("dog eat")
}
func (c cat) say(name string)  {
	fmt.Println("I'm a cat .. name is", name)
}

// 空接口
func null_inter(in interface{}) {
	v,ok := in.(string)
	if !ok {
		fmt.Println("非 string")
	}else {
		fmt.Println("是string  ", v)
	}
}
// 空接口switch case
func null_switch_interface(in interface{}) {
	switch v := in.(type) {
	case string:
		fmt.Println("type is string", v)
	case int:
		fmt.Println("type is int", v)
	case say:
		fmt.Println("type is say")
	}
}