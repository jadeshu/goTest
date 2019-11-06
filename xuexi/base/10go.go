/**
 * @Author: jadeshu
 * @Description:
 * @File:  10go
 * @Version: 1.0.0
 * @Date: 2019/11/6 18:42
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(n int) {
			for {
				fmt.Printf("当前携程：%d\n", n)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
