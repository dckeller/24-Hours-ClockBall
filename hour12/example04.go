package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(3 * time.Second)
	for {
		c <- "ping"
		test := <-t.C
		fmt.Println(test)
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
