package main

import "fmt"

//type person struct {
//	age  int
//	name string
//}

func ret() *Person {
	p := &Person{12, "jadeshu"}
	// p的值 因为是指针
	fmt.Printf("&person:%p\n", p)
	// p的地址,p是变量有地址
	fmt.Printf("&person:%p\n", &p)

	return p
}

func ret2() Person {
	p := Person{12, "jadeshu"}
	fmt.Printf("person:%T ,%p\n", p, &p)
	//fmt.Printf("%v\n",&p)

	return p
}

func main() {
	per1 := ret()
	fmt.Printf("%T ,%p ,%p\n", per1, per1, &per1)

	per := ret2()
	b := per == *per1
	fmt.Println(b)
	fmt.Printf("%T ,%v ,%p\n", per, per, &per)
}
