package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l)
	l.Back()
	l.PushBack("back")
	l.PushFront("frist")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
