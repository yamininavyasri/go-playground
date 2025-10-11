/*
Channels are typed conduits that allow goroutines to communicate with each other.
You can send a value into a channel from one goroutine and receive it from another.
Channels help to synchronize execution and share data safely between goroutines.

Brief explanation:
ch := make(chan string) → Create a channel for string messages.
go worker(ch) → Launch a goroutine that does some work.
Inside worker, ch <- "Task completed" sends the message to the channel.
message := <-ch → Main goroutine blocks until it receives a value from the channel.
Output shows how main waits for worker without explicit locks.
*/
package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	fmt.Println("Worker started")
	time.Sleep(2 * time.Second)
	ch <- "Task completed" // send value to channel
}

func main() {
	ch := make(chan string) // create channel

	go worker(ch) // start worker goroutine

	fmt.Println("Waiting for worker...")
	message := <-ch // receive message from channel
	fmt.Println("Received:", message)
}
