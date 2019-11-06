package main

import (
	"log"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	type iv struct {
		int
	}

	i := iv{}
	log.Printf("Sending user email to %s", u.name)
	println(i)
}

func main() {
	u := user{"Bill", "bill@email.com"}

	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
