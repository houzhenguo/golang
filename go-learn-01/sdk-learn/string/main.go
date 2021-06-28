package split

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// string 练习
	// Fields 用一个或多个连续的空格分割字符串s,返回子字符串数组slice
	fmt.Printf("Fields are : %q", strings.Fields(" foo bar baz  "))

	// split 不带 sep
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	// split 带 sep
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))

	fmt.Println("===========Split ======== 官方")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // 注意空格
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println("============前缀/后缀============")
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", "")) // 注意空字符串的处理
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
	fmt.Println("============ index========")

	// 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
	han := func(c rune) bool {
		return unicode.Is(unicode.Han, c) // 汉字
	}
	// 当 字符串满足这个函数当时候找到索引
	h := func(c rune) bool {
		// fmt.Println("c rune is", c)
		return c == 'l' // ascii 表
	}
	fmt.Println(strings.IndexFunc("Hello, world", han))
	fmt.Println(strings.IndexFunc("Hello, 世界", han)) // 汉字开头的index
	fmt.Println(strings.IndexFunc("Hello, world", h))
	// test make slice
	sl := make([]byte, 2, 4)
	fmt.Println("sl length is ", len(sl), "---", cap(sl))
	// 重复n 次
	fmt.Println("ba" + strings.Repeat("na", 2))
	// string 中Builder 测试
	var b strings.Builder
	b.Grow(10)                   // 判断 cap - len > n 否则扩容，至少n ,2 * len(b) + n
	b.WriteString("hello world") // 看到 string中频繁使用这个类，所以想试一下
	fmt.Println("ttt ", b.String())
	// 字符串替换 map 函数
	rp := func(s rune) rune {
		if s >= 'a' && s <= 'z' {
			s -= 32
			return s
		}
		if s >= 'A' && s <= 'Z' {
			s += 32
			return s
		}
		if han(s) {
			return '!'
		}
		return 0
	}
	afc  := strings.Map(rp, "hello 你好 WORLD 中国") //  相当于替换大小写吧
	fmt.Println("afc is", afc) // afc is HELLO!!world!!

	// Replace 替换函数
	rp1 := strings.Replace("hello hou zh", "h", "m", 2) // 可以使用 -1
	fmt.Println("rp1 is", rp1)
	// Trim 函数
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))

	// 字符串转 int
	// n1 is  127 error is strconv.ParseInt: parsing "128": value out of range
	n1, err := strconv.ParseInt("128", 10, 8) // n1 代表能存储的最大值， err返回具体的错误
	fmt.Println("n1 is ", n1, "error is", err)
}
