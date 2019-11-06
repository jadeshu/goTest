package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle hello")
		fmt.Fprintf(w, "hello")
	})

	err := http.ListenAndServe("127.0.0.1:6060", nil)
	if err != nil {
		fmt.Println("listen fail!")
	}

}

//func Hello(w http.ResponseWriter,r *http.Request)  {
//	fmt.Println("handle hello")
//	fmt.Fprintf(w,"hello")
//}
