package main

import "fmt"

// `1 1 2 3 5 8 13 21 .......
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 函数递归
func sum1(v int) int {
	if v <= 1 {
		return 1
	}
	return v + sum1(v-1)
}

// 函数式1
func sum2(val int) func() int {
	sum := 0
	return func() int {
		sum += val
		return sum
	}
}

// 函数式2
func sum3() func(val int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}

//===正统式函数编程===
//*不可变性：不能有状态，只有常量和函数
//*函数只能有一个参数
type iSum func(int) (int, iSum)

func sum4(base int) iSum {
	return func(val int) (i int, sum iSum) {
		return base + val, sum4(base + val)
	}
}
func main() {
	//// `1 1 2 3 5 8 13 21 .......
	//f := fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	// 函数递归
	//fmt.Println(sum(100))

	// 函数式1
	//f2 := sum2(10)
	//fmt.Println(f2())
	//fmt.Println(f2())

	// 函数式2
	//f3 := sum3()
	//for i := 1; i < 10; i++ {
	//	fmt.Printf("0 + ...... + %d = %d\n", i, f3(i))
	//}

	//===正统式函数编程===
	f4 := sum4(0)
	for i := 1; i < 10; i++ {
		var s int
		s, f4 = f4(i)
		fmt.Printf("0 + ...... + %d = %d\n", i, s)
	}

}
