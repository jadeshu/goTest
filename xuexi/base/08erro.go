package main

import (
	"fmt"
	"os"
)

func div(a, b int) (res int, err error) {
	defer func() { // 发生错误时恢复
		if e := recover(); e != nil {
			err = fmt.Errorf("err: %v", e)
		}
	}()

	if b == 0 {
		panic("分母不能为0") // 在这宕机 后面程序不执行了
	} else {
		res = a / b
	}
	return
}

func main() {
	res, err := div(10, 0)
	fmt.Println(err.Error())
	fmt.Println(res)
	os.Create()
}
