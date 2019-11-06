package main

import "fmt"

//const v,iv = "1212",12
const t int = 23
const (
	v  = "1212"
	iv = 12
)

var (
	yy int
	zz int
)

func main() {
	const mx = 10
	fmt.Println(iv)
	fmt.Printf("%#v , %d\n", v, mx)

	i, j := 10, 20
	val, _ := i, j
	fmt.Println(val)
	fmt.Println(j)

	cp := 12 + 3i
	fmt.Printf("%T\n", cp)
	fmt.Println(real(cp), imag(cp))
	var s string
	fmt.Scanf("%s", &s)
	fmt.Scan(&s)
	fmt.Println(s)
}
