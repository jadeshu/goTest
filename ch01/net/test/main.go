package main

import (
	"jadeshu.cn/studyGo/ch01/net/cl"
)

func main() {
	cl.Client()
}

//func client() {
//	conn, err := net.Dial("tcp", ":8080")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if _, err := conn.Write([]byte("helloworld")); err != nil {
//		log.Fatal(err)
//	}
//	if err := conn.Close(); err != nil {
//		log.Fatal(err)
//	}
//}
