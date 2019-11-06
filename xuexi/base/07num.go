package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

// 产生一个随机4位数
func randNum4(num *int) {
	rand.Seed(time.Now().UnixNano())
	var p int
	for {
		p = rand.Intn(10000)
		if p > 999 {
			break
		}
	}

	*num = p
}

// 获取4位数的每一位
func getNum4(s []int, num int) {
	s[0] = num / 1000
	s[1] = (num % 1000) / 100
	s[2] = (num % 100) / 10
	s[3] = num % 10
}

// 循环输入4位数与随机数进行比较
func do(key []int) {
	var num int
	for {
		for {
			fmt.Println("请输入一个数")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入合法的四位数")
		}
		s := make([]int, 4)
		getNum4(s, num)

		n := 0
		for i := 0; i < 4; i++ {
			if s[i] > key[i] {
				fmt.Printf("第%d位猜大了\n", i+1)
			} else if s[i] < key[i] {
				fmt.Printf("第%d位猜小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}

		if n == 4 {
			break
		}
	}

}

func main() {
	var sradnum int
	fmt.Println(unsafe.Sizeof(sradnum))
	randNum4(&sradnum)
	//fmt.Println(sradnum)
	// 保存4位数
	s := make([]int, 4)
	getNum4(s, sradnum)
	fmt.Println(s)

	do(s)
}
