package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	age   int
	name  string
	scorp map[string]int
}

func main() {
	var st student
	fmt.Println(unsafe.Offsetof(st.scorp))
	fmt.Println(unsafe.Alignof(st.scorp))
	// 基本数据类型
	fmt.Println("======基本数据类型======")
	var iNum8 int8
	fmt.Println(unsafe.Sizeof(iNum8)) // 1
	var iNum16 int16
	fmt.Println(unsafe.Sizeof(iNum16)) // 2
	var iNum32 int32
	fmt.Println(unsafe.Sizeof(iNum32)) // 4
	var iNum64 int64
	fmt.Println(unsafe.Sizeof(iNum64)) // 8
	var uiNum32 uint32
	fmt.Println(unsafe.Sizeof(uiNum32)) // 4
	var re rune
	fmt.Println(unsafe.Sizeof(re)) // 4
	var iNum int
	fmt.Println(unsafe.Sizeof(iNum)) // 8
	var fNum32 float32
	fmt.Println(unsafe.Sizeof(fNum32)) // 4
	var fNum64 float64
	fmt.Println(unsafe.Sizeof(fNum64)) // 8
	var str string
	fmt.Println(unsafe.Sizeof(str)) // 16
	var arr [10]int
	fmt.Println(unsafe.Sizeof(arr)) // 80=10*8

	fmt.Println("======其他类型======")
	// 其他类型
	// 指针类型
	var ptr *int
	fmt.Println(unsafe.Sizeof(ptr)) // 8
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8

	// 切片
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	// map
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	// 通道
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	// 函数
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	// 接口
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16
}
