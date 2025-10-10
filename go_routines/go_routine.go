/*A goroutine is a lightweight thread managed by the Go runtime.

In simple terms:
A goroutine allows you to run functions concurrently (at the same time) — without manually dealing with threads or complex synchronization.
Every Go program starts with one main goroutine (the main() function).
If the main goroutine finishes, all other goroutines stop automatically, even if they haven’t finished their work yet.

That’s why we often use time.Sleep(), sync.WaitGroup, or channels to keep the main goroutine alive until others finish.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func say(word string) {
	fmt.Println(word)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signal done when finished
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	go say("hello from go routine")
	time.Sleep(time.Second)
	fmt.Println("hello from main")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // wait for all goroutines
	fmt.Println("All workers finished")
}
