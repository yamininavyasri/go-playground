// context.go (no tag)
/*
The context package is used to carry deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.It helps you control how long a function or goroutine should run — and allows you to cancel operations cleanly when they’re no longer needed (for example, when a request times out or a user disconnects).
Why we need context

Imagine you have a web server that starts multiple goroutines for each request:

Each goroutine queries a database,

Fetches data from an external API,

And does some processing.

If the client cancels the request, you don’t want those goroutines to keep running unnecessarily.
This is where context helps — it signals them to stop gracefully.

Creating Contexts

You usually start with a base context and then create derived contexts from it.
1. context.Background()
The root context.
Used in main functions, initialization, or tests.
2. context.TODO()
Placeholder when you’re not sure which context to use yet.

| Function                                | Purpose                        |
| --------------------------------------- | ------------------------------ |
| `context.WithCancel(parent)`            | Manually cancel operations     |
| `context.WithTimeout(parent, duration)` | Auto-cancel after a timeout    |
| `context.WithDeadline(parent, time)`    | Auto-cancel at a specific time |
| `context.WithValue(parent, key, value)` | Pass request-scoped values     |

Note: Context values are meant for request-scoped data (like trace IDs, auth tokens) — not for passing optional parameters.
| Method       | Description                                                                         |
| ------------ | ----------------------------------------------------------------------------------- |
| `Done()`     | Returns a channel that closes when context is cancelled                             |
| `Err()`      | Returns error after cancellation (`context.Canceled` or `context.DeadlineExceeded`) |
| `Deadline()` | Returns when the context will be cancelled                                          |
| `Value(key)` | Gets value associated with the key                                                  |

summary:
| Feature         | Description                                                |
| --------------- | ---------------------------------------------------------- |
| **Purpose**     | Manage cancellation, timeout, and values across goroutines |
| **Key funcs**   | `WithCancel`, `WithTimeout`, `WithValue`                   |
| **Common uses** | HTTP servers, DB queries, background jobs, API calls       |
| **Benefit**     | Prevent goroutine leaks and handle timeouts gracefully     |
*/

package main

import (
	"context"
	"fmt"
	"time"
)

/*
Explanation:
Program starts and creates a cancellable context.
Goroutine sleeps for 5 seconds, then calls cancel().
The select statement is waiting:
After 5 seconds, <-ctx.Done() is closed.
Prints:
Context cancelled
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	/*
		select waits for multiple channel operations simultaneously:
		<-ctx.Done() → triggered if the context is cancelled.
		<-time.After(10 * time.Second) → triggered if 10 seconds pass.
		Whichever happens first executes.
	*/
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	case <-time.After(10 * time.Second):
		fmt.Println("Operation completed")
		//default:
		//	fmt.Println("Operation not completed")
	}
}
