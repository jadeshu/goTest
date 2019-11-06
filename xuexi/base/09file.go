package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 新建文件
	file, e := os.Create("ceshi.txt")
	if e != nil { // 发送错误时
		fmt.Println("文件失败！")
	}
	// 写入字符串
	file.WriteString("this is a test!\n")

	// 可以写入字符串和二进制文件
	file.Write([]byte("jadeshu\n"))
	file.Write([]byte("jajjaj"))

	// 关闭文件
	file.Close()

	// 打开文件  只读模式   其他模式可以用OpenFile指定模式打开
	open, e := os.Open("ceshi.txt")
	if e != nil {
		fmt.Println("打开文件失败！")
	}
	//s := make([]byte, 1024)
	//// 读取文件
	//n, err := open.Read(s)
	//// 文件读取出错同时没有到文件结尾
	//if err != nil && err != io.EOF {
	//	fmt.Println("erro=", err)
	//}
	//fmt.Println("=====输出读取的字节========")
	//fmt.Println(n)
	//fmt.Println("=====输出二进制========")
	//fmt.Println(s[:n])
	//fmt.Println("=====输出字符串========")
	//fmt.Println(string(s[:n]))
	//fmt.Println("=====输出字符串========")
	////fmt.Println(s)
	//fmt.Printf("%s", string(s))

	// 借助bufio读取行
	// 新建缓冲区
	reader := bufio.NewReader(open)
	line, _, _ := reader.ReadLine()
	fmt.Println("=====bufio========")
	fmt.Println(string(line))

	line, _, _ = reader.ReadLine()
	fmt.Println("=====bufio========")
	fmt.Println(string(line))

	open.Close()
}
