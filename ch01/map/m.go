package main

import "fmt"

func main() {
	delete()
	//// 随机数种子
	//rand.Seed(time.Now().UnixNano())
	//
	//scoreMap := make(map[string]int, 200)
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("student%02d", i)
	//	scoreMap[key] = rand.Intn(100)
	//}
	//
	//fmt.Println(scoreMap)
	//
	//// 排序map的key
	//// 按字符串排序
	//keys := make([]string, 0, 100)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	//
	//sort.Strings(keys)
	//// 打印
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}
	var mapSlice = make([]map[string]string, 1, 3)
	for index, value := range mapSlice {
		fmt.Printf("===index:%d value:%v\n", index, value)
	}

	cur := map[string]string{"name": "hahah", "1212": "3333"}
	mapSlice = append(mapSlice, cur)
	for index, value := range mapSlice {
		fmt.Printf("!!!index:%d value:%v\n", index, value)
	}

	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "jadeshu"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "北京"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}

	//a := make([]int, 10)
	//fmt.Println(a)      //[0 0]
	//a[1] = 20
	//fmt.Println(a)
	//fmt.Println(len(a)) //2
	//fmt.Println(cap(a)) //10

	// 端口：【address,ip】
	//info := make(map[int]map[string]string, 3)
	//info[8080] = make(map[string]string,4)
	//info[8080]["address"] = "上海"
	//info[8080]["ip"] = "127.0.0.1"
	//info[6060] = make(map[string]string,4)
	//info[6060]["address"] = "北京"
	//info[6060]["ip"] = "192.168.0.101"
	//fmt.Println(info)
	//
	//testInfo:= &map[int]map[string]string{
	//	6000:{
	//		"address":"香港",
	//		"ip":"127.0.0.1",
	//	},
	//	7000:{
	//		"address":"武汉",
	//		"ip":"127.0.0.1",
	//	},
	//}
	//
	//fmt.Println(testInfo)

}
