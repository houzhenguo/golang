package main

import (
	"fmt"
	"sort"
)

// sort 练习
// 定义一个Person ，可以根据Person 的 年龄 / 身高 / 成绩进行动态的排序-> 如果
// 可以传入自定义的排序函数
func main() {
	arr := PersonArr{
		{"zs", 19,65.7},
		{"ls", 25,96.7},
		{"wu", 23,60.1},
	}
	age := func(p1, p2 *Person) bool {
		return p2.age > p1.age
	}
	//score := func(p1, p2 *Person) bool {
	//	return p1.score > p2.score
	//}
	SortFunc(age).Sort(arr)
	fmt.Println(arr)
	SortFunc(func(p1, p2 *Person) bool {
		return p1.score < p2.score
	}).Sort(arr)
	fmt.Println(arr)
}

// Person 定义struct
type Person struct {
	name  string
	age   int
	score float32
}

// PersonArr 定义数组
type PersonArr []Person
type SortFunc func(p1, p2 *Person) bool
type PersonSorter struct {
	persons []Person
	// 比较函数
	sort SortFunc
}

func (by SortFunc) Sort(persons []Person) () {
	ps := &PersonSorter{
		persons: persons,
		sort:    by,
	}
	sort.Sort(ps)
}

func (arr *PersonSorter) Len() int {
	return len(arr.persons)
}

func (arr *PersonSorter) Swap(i, j int) {
	arr.persons[j], arr.persons[i] = arr.persons[i], arr.persons[j]
}
func (arr *PersonSorter) Less(i, j int) bool {
	return arr.sort(&arr.persons[i], &arr.persons[j])
}
