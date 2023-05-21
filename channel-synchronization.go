package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <-  true
}

func main(){
	done := make(chan bool, 1)

	go worker(done)

	<-done //Block until we recieve a notification from the worker on the channel
}