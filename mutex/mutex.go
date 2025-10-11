/*
Mutex (Mutual Exclusion Lock) ensures that only one goroutine accesses a shared resource (like a variable or map) at a time.
Go provides this in the sync package.
Why We Need It:
When multiple goroutines try to read/write to the same variable at the same time, it causes a race condition.

mu.Lock() → locks the shared resource.
mu.Unlock() → unlocks it so others can access.
Ensures no two goroutines increment count at the same time.
sync.WaitGroup ensures the program waits until all goroutines finish.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Launch 10 goroutines that increment the counter 1000 times each
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// Uncomment this line to see incorrect results (no synchronization)
				//counter++

				// Correct way: protect access with a mutex
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
