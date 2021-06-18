package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello reflect")
	var a float32 = 3.8
	reflectType(a)
	var b int64 = 88
	reflectType(b)
	d := dog{
		name:"www",
	}
	reflectType(d)
	reflectType(&d) // 需要区分指针的时候使用kind  v type is  and kind is ptr

	// reflect value
	reflectValue(int64(11))
	reflectValue(float32(3.3))
	// 通过反射修改值
	var chv int64 = 22
	reflectSetValue1(&chv)
	fmt.Println("chv is", chv)
}
// type 感觉就是到倒数第二层
// kind 就是本质
func reflectType(x interface{}) {
	v:= reflect.TypeOf(x)
	fmt.Printf("v type is %v and kind is %v\n", v.Name(), v.Kind())
}
type dog struct {
	name string
}
// 通过反射获取值
func reflectValue(x interface{}) {
	v:= reflect.ValueOf(x)
	k:= v.Kind()
	switch k {
	case reflect.Int64:
		v1:= v.Int()
		fmt.Printf("v1 is 原始值 %d ，强制转换 long之后的值%d\n", v1, v1)
	case reflect.Float32:
		v2 := v.Float()
		fmt.Printf("v2 is 原始值 %f ，强制转换 long之后的值%f\n", v2, float32(v2))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	//if v.Kind() == reflect.Int64 {
	//	v.SetInt(200) //修改的是副本，reflect包会引发panic
	//}
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(222) //  外面传递的是 &指针
	}
}
