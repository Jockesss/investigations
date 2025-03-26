package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: Sending message...")
		messages <- "ping"              // (1)
		fmt.Println("B: Message sent!") // (2)
	}()

	fmt.Println("A: Doing some work...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-messages //  (3)

	fmt.Println("A: Messege received!")
	time.Sleep(100 * time.Millisecond)
}
