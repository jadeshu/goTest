package main

import "fmt"

// 声明变量
// var name string
// var iB int

// 批量声明
var (
	name   string //""
	age    int    //0
	isFlag bool   //false
)

/*
*注意：
*声明了全局变量可以不使用，编译器不会报错(因为别人用这个包可能会使用)
*声明了局部变量必须使用，否则编译器报错
*
*go fmt xxx.go  命令可以格式化代码
 */

//str4 := "cuowu"  //出错

func main() {
	// 1.赋值
	name = "jadeshu"
	age = 20
	isFlag = true

	fmt.Print(name)
	fmt.Println()
	fmt.Println(isFlag)
	fmt.Printf("age:%d", age)

	// 2.变量声明并赋值
	var str string = "hehehe"
	fmt.Println(str)
	// 自动根据数值进行推导类型
	var str2 = "hahhahh"
	fmt.Println(str2)

	// 3.简单变量声明并赋值
	// 只能在局部变量进行短声明赋值 用的较多
	// 不能重复声明，且必须声明时赋值
	str3 := "heaaaa"

	// 这个情况下可以，只要简单赋值语句里有一个变量是新声明的就可以
	// 同时还能改变原有变量的值
	str3, str4 := "1212", "sdsdsd"
	fmt.Println(str3)
	fmt.Println(str4)
	// 4.匿名变量 用一个下划线表示
	// 当忽略某个值时就用它
	// 它不会占用内存，也不存在重复声明
	// GO语言中函数可以返回多个值，这时候很有用
	a, _ := ret() // 忽略返回的y值
	fmt.Println(a)
}

func ret() (x, y int) {
	return 10, -1
}
