package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	start := time.Now()

	fmt.Println("Worker started and waiting for cancellation...")
	// This select will **block indefinitely** until ctx.Done() is closed
	select {
	case <-ctx.Done():
		elapsed := time.Since(start).Seconds()
		fmt.Printf("Worker stopped due to cancellation after %.2f seconds\n", elapsed)
	}
	elapsed := time.Since(start).Seconds()
	fmt.Printf("Worker finished all tasks in %.2f seconds\n", elapsed)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second) // simulate some external event
		cancel()                    // cancel the context after 3 seconds
	}()

	worker(ctx)
	fmt.Println("Main function exits")
}
