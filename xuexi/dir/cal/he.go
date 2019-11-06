package cal

import "fmt"

var Std = 16

type student struct {
	age  int
	name string
}

const Str = "常量"

func Move() {
	fmt.Println("cal::Move()")
}
