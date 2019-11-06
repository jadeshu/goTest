package main

import "fmt"

// 常量在运行期间不允许值改变
// 一个常量声明
const PI = 3.14

// 多个常量一起声明
const (
	DIR = 2 * PI
)

// t1 t2是100 t3 t4是200
// 同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	t1 = 100
	t2
	t3 = 200
	t4
)

// iota是go语言的常量计数器
// const中每新增一行常量声明将使iota计数一次
// iota可以理解为const语句块中的行索引，
// 使用iota能简化定义，在定义枚举时很有用
const (
	m1 = iota
	m2
	m3
)

// 增强理解iota
const (
	n1 = 10
	n2
	n3 = iota
	n4
)

// 增强理解iota
// 多个
const (
	s1, s2 = iota + 1, iota + 2
	s3, s4 = iota + 1, iota + 2
)

// 练习
// const (
// 	_ = iota
// 	KB = 1 << (iota*10)
// 	MB = 1 << (iota*10)
// 	GB = 1 << (iota*10)
// 	TB = 1 << (iota*10)
// 	PB = 1 << (iota*10)
// )

func main() {
	fmt.Println(DIR)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println("============")

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println("============")

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println("============")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println("============")
}
