package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func (p person) testPer() {
	fmt.Println("person::testPer()")
}

// 注意两者区别
type man person // 另外类型
//type man = person   // 别名

func (m man) getAge() int {
	return m.age
}

func fn() func() int {
	sum := 0
	return func() int {
		sum++
		return sum * sum
	}
}

func main() {
	m := man{10, "hahha"}
	per := person{20, "jade"}
	fmt.Printf("%T  --%T\n", m, per)
	fmt.Println(m.getAge())
	fmt.Println(m.name)
	//per.getAge()
	//fmt.Println(per.getAge())

	// 匿名函数
	//val := 20
	//func(i int) {
	//	fmt.Println(i)
	//	i = 300
	//	fmt.Println(i)
	//}(val)
	//fmt.Println(val)

	//====闭包============
	//fn := func() func() int {
	//	num := 0
	//	return func() int {
	//		num++
	//		return num * num
	//	}
	//}
	f := fn()

	fmt.Println("====")
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println() // 1
	//fmt.Println(fn()()) // 1
	//fmt.Println(fn()()) // 1
	//fmt.Println(fn()()) // 1

}

//package main
//import (
//	"fmt"
//	"reflect"
//)
//// 定义商标结构
//type Brand struct {
//	age int
//}
//// 为商标结构添加Show()方法
//func (t Brand) Show() {
//	fmt.Println("func (t Brand) Show()")
//}
//// 为Brand定义一个别名FakeBrand
//type FakeBrand = Brand
//// 定义车辆结构
//type Vehicle struct {
//	e func() int
//	// 嵌入两个结构
//	FakeBrand
//	Brand
//}
//
//func add() int  {
//	fmt.Println("add func")
//	return  10
//}
//
//func main() {
//	// 声明变量a为车辆类型
//	var a Vehicle
//	a.e = add
//	//
//	//a.e()
//	//fmt.Printf("FieldType: %T\n", a.e)
//	// 指定调用FakeBrand的Show
//	a.Brand.age = 20
//	a.FakeBrand.age = 100
//
//	// 函数调用的是一个
//	a.FakeBrand.Show()
//	a.Brand.Show()
//
//	// 对象不同
//	fmt.Println(a.Brand.age)
//	fmt.Println(a.FakeBrand.age)
//
//	// 取a的类型反射对象
//	ta := reflect.TypeOf(a)
//	// 遍历a的所有成员
//	for i := 0; i < ta.NumField(); i++ {
//		// a的成员信息
//		f := ta.Field(i)
//		// 打印成员的字段名和类型
//		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
//	}
//}
