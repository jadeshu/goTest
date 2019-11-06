package cl

import (
	"fmt"
	"log"
	"net"
)

func init() {
	fmt.Printf("cl ---模块初始化!")
}

func Client() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write([]byte("helloworld")); err != nil {
		log.Fatal(err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}
