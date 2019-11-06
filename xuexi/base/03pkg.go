package main

////导入单个包
//import "fmt"
//
//func main() {
//	fmt.Println("导入包操作案例!")
//}

//导入多个包方式
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("导入包操作案例!")
	fmt.Println(os.Args)
}

//// 忽略包
//// 应用场景 去掉此包
//import _ "fmt"
//
//func main() {
//	//fmt.Println("121212")
//}

////.操作
//import . "fmt"
//func main() {
//	Println("点操作包的方式")
//}

////给包起别名
//import xx "fmt"
//func main() {
//	xx.Println("给包起别名操作")
//}
