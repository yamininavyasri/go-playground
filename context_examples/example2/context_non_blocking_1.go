/*
How it works

The worker goroutine loops 10 times.

Inside the loop, the select does:

<-ctx.Done() → checks if context is cancelled.

default → runs immediately if context is not cancelled (non-blocking).

After 2 seconds, the anonymous goroutine calls cancel().

Worker detects cancellation without waiting, prints a stop message, and exits.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	start := time.Now() // record start time

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			elapsed := time.Since(start).Seconds() // elapsed seconds
			fmt.Printf("Worker stopped due to cancellation after %.2f seconds\n", elapsed)
			return
		default:
			fmt.Println("Worker processing item", i)
			time.Sleep(500 * time.Millisecond) // simulate work
		}
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("Worker finished all tasks in %.2f seconds\n", elapsed)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("Sleeping for 2 seconds")
		time.Sleep(2 * time.Second) // simulate external stop signal
		cancel()
	}()

	worker(ctx)
	fmt.Println("Main function exits")
}
