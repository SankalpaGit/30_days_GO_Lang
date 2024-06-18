package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan<- string) {
	time.Sleep(1 * time.Second)
	msg := "Hello from Testing User !"
	ch <- msg
}

func receiveMessage(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	receiveMessage(ch)
}
