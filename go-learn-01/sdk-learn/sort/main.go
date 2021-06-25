package main

import (
	"fmt"
	"sort"
)

func main1() {
	s1 := student{"zhangsan",18}
	s2 := student{"lisi",20}
	s3 := student{"wangwu",16}
	students := studentArr{s1, s2, s3} //
	sort.Sort(students)
	fmt.Println(students)
}

type student struct {
	name string
	age  uint
}

type studentArr []student

func (students studentArr) Len() int {
	return len(students)
}
func (students studentArr) Less(i, j int) bool {
	return students[i].age > students[j].age
}
func (students studentArr)Swap(i, j int)  {
	students[j],students[i] = students[i],students[j]
}
