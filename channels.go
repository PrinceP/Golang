package main

import "fmt"

func main() {

	messages:= make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	go func() { messages <- "pong" }()

	msg = <-messages
	fmt.Println(msg)

	messages_2 := make(chan string, 2)
	messages_2 <- "buf1"
	messages_2 <- "buf2"

	fmt.Println(<-messages_2)
    fmt.Println(<-messages_2)
	

}