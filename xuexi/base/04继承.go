package main

import "fmt"

type person struct {
	age  int
	name string
}

// 方法
func (p *person) display() {
	fmt.Println("person:", p.age, p.name)
	fmt.Println(p)
}

type stduent struct {
	*person
	id int
}

// 在stduent中重写person的display()方法，进行覆盖person:display()
func (s *stduent) display() {
	fmt.Println("stduent:", s.age, s.name, s.id)
	fmt.Println(s)
}

func main() {
	s := stduent{&person{12, "jadeshu"}, 0001}
	// 调用基类函数
	//s.person.display()
	// 调用子类函数
	s.display()
}
