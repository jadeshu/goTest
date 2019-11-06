package main

import "fmt"

func main() {
	//s1 := make([]int, 5, 6)
	//fmt.Println(s1)
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))

	// 数组
	s2 := [3]int{1, 2, 3}
	s3 := s2 // 赋值 ，内存不同
	fmt.Println(s2, s3)
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%p\n", &s3)
	fmt.Println(s2 == s3) // 注意区别，值比较
	s2[0] = 1000
	fmt.Println(s2, s3) //[1000 2 3] [1 2 3]

	fmt.Println(&s2 == &s3) // 注意区别，取得的是地址
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
	fmt.Println("==================")

	// 切片
	s4 := []int{1, 2, 3}
	s5 := s4 // 引用s4里的值指向一个地方
	fmt.Println(s4, s5)
	s4[0] = 1000
	fmt.Println(s4, s5) //[1000 2 3] [1000 2 3]

	fmt.Println(&s4 == &s5) // 注意区别，取得的是变量的地址
	fmt.Printf("%p\n", s4)
	fmt.Printf("%p\n", s5)

	fmt.Println("===============")
	//test()

	fmt.Printf("\n")

	// make切片
	mkSlice := make([]int, 5)
	fmt.Println(mkSlice)

	// 追加到了末尾
	mkSlice = append(mkSlice, 10)
	fmt.Println(mkSlice)

	// make切片 并追加到切片头部
	// 那么需要把长度设为0，即没有切片内容
	//  长度为0，容量为5
	mkSlice2 := make([]int, 0, 5)
	fmt.Println(mkSlice2)

	// 追加到了末尾
	mkSlice2 = append(mkSlice2, 10)
	fmt.Println(mkSlice2)

}

func test() {
	slice0 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("\n~~~~~~元素遍历~~~~~~")
	for _, ele := range slice0 {
		fmt.Print(ele, "-")
		ele = "7" // ele值改变不了切片中的值
	}

	fmt.Println("\n========")
	fmt.Println(slice0)

	fmt.Println("\n~~~~~~索引遍历~~~~~~")
	for index := range slice0 {
		fmt.Print(slice0[index], " ")
	}
	fmt.Println("\n~~~~~~元素索引共同使用~~~~~~")
	for index, ele := range slice0 {
		fmt.Print(ele, "-", slice0[index], " ")
	}
	fmt.Println("\n~~~~~~修改~~~~~~")
	for index := range slice0 {
		slice0[index] = "9"
	}
	fmt.Println(slice0)
}
