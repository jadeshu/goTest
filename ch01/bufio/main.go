package main

import "fmt"

func main() {
	//// 准备从标准输入读取数据
	//inReader := bytes.NewReader(os.Stdin)
	//fmt.Println("请输入您的名字：")
	//
	//// 读取数据直到\n为止
	//input, err := inReader.ReadString('\n')
	//if err != nil {
	//	fmt.Printf("错误：%s\n", err)
	//	os.Exit(1)
	//} else {
	//	// 用切片删除最后的\n
	//	name := input[:len(input)-1]
	//	fmt.Printf("hello %s!\n", name)
	//}
	//
	//for {
	//	input, err = inReader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("An error occurred: %s\n", err)
	//		continue
	//	}
	//	input = input[:len(input)-1]
	//	// 全部转换为小写。
	//	input = strings.ToLower(input)
	//	switch input {
	//	case "":
	//		continue
	//	case "nothing", "bye":
	//		fmt.Println("Bye!")
	//		// 正常退出。
	//		os.Exit(0)
	//	default:
	//		fmt.Println("Sorry, I didn't catch you.")
	//	}
	//}

	arr := []int{21, 32, 12, 33, 34, 34, 87, 5}
	var n = len(arr)
	fmt.Println("--------没排序前--------\n", arr)
	for i := 0; i < n-1; i++ {
		fmt.Println("--------第", i+1, "次冒泡--------")
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				//t := arr[i]
				//arr[i] = arr[j]
				//arr[j] = t
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("--------最终结果--------\n", arr)

}
