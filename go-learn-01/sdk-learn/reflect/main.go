package main

import (
	"fmt"
	"reflect"
)

type newInt int64

func main() {
	var a float32 = 2.34
	reflectType(a)
	var b = "hello world"
	reflectType(b)
	var c = &struct {
		string
	}{"zs"}
	reflectType(c)
	var ni newInt = 78 // type newInt, kind = ptr
	reflectType(ni)
}

// kind 我理解是更底层的，一般用在 判断 struct / ptr 这方面
func reflectType(x interface{}) {
	// 通过reflectType 获取类型
	v := reflect.TypeOf(x)
	fmt.Printf("type :%v , kind is:%v \n", v.Name(), v.Kind())
	reflectValue(int32(1))
	reflectValue(1.2)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int32:
		var a = int32(v.Int())
		fmt.Printf("a is %d and type is %T\n", a, a)
	case reflect.Float64:
		var b = v.Float()
		fmt.Printf("a is %v and type is %T\n", b, b)
	}
	// Set int 就是设置值
	var d1 = int32(23)
	reflectSetValue1(&d1)
	fmt.Println("d1 is ", d1)
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	if k == reflect.Int8 {
		v.SetInt(12)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	if k == reflect.Int32 {
		v.Elem().SetInt(33)
	}
}