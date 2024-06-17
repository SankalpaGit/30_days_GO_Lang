package main

import (
	"fmt"
	"time"
)

func printMessage(message string, done chan struct{}) {
	fmt.Println(message)
	time.Sleep(1 * time.Second)
	close(done)
}

func main() {
	done := make([]chan struct{}, 9)

	// Initialize channels
	for i := range done {
		done[i] = make(chan struct{})
	}

	// Start goroutines
	go printMessage("First message", done[0])

	go func() {
		printMessage("Second message", done[1])
	}()
	go func() {
		<-done[1]
		printMessage("Third message", done[2])
	}()
	go func() {
		<-done[2]
		printMessage("Fourth message", done[3])
	}()
	go func() {
		<-done[3]
		printMessage("Fifth message", done[4])
	}()
	go func() {
		<-done[4]
		printMessage("Sixth message", done[5])
	}()
	go func() {
		<-done[5]
		printMessage("Seventh message", done[6])
	}()
	go func() {
		<-done[6]
		printMessage("Eighth message", done[7])
	}()
	go func() {
		<-done[7]
		printMessage("Ninth message", done[8])
	}()

	<-done[8]
	fmt.Println("Main function exiting")
}
