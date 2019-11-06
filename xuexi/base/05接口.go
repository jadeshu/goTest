package main

import "fmt"

type HUhi interface {
	sayHi()
}

type student1 struct {
}

func (s student1) sayHi() {
	fmt.Println("student hi....")
}

type dog struct {
}

func (d dog) sayHi() {
	fmt.Println("dog hi....")
}

func whoHi(uhi HUhi) {
	uhi.sayHi()
}

func main() {
	s := student1{}
	d := dog{}
	whoHi(s)
	whoHi(d)

	defer fmt.Println("test")
	defer fmt.Println("defer test")

}
