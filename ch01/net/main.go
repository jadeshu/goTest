package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	server()
	//client()
}

func server() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		client, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			io.Copy(os.Stdout, c)
			c.Close()
		}(client)
	}
}
