/*
| Concept     | Behavior                                                   |
| ----------- | ---------------------------------------------------------- |
| `defer`     | Executes after surrounding function returns                |
| `panic`     | Stops normal execution immediately                         |
| `recover`   | Catches panic inside a deferred function                   |
| Combination | Defer + Recover allows **clean exit** and resource cleanup |

After a panic in a function:
Normal execution stops immediately at the point of the panic.
No subsequent lines in that function are executed after the panic, unless they are deferred.

Deferred functions still run before the function exits.
If recover() is used in a deferred function, you can catch the panic and continue execution in the caller.
*/

package main

import "fmt"

func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Dividing", a, "by", b)
	result := a / b // will panic if b == 0
	fmt.Println("Result:", result)
}

/*
for example 2:
Your program panicked in the demo() function.
Go printed a stack trace, showing exactly where the panic occurred and the call path that led to it.
Since you didn’t recover from the panic, the program exited immediately with status 2.
*/
func demo() {
	defer fmt.Println("Deferred runs even after panic")
	fmt.Println("Before panic")
	panic("Panic occurred")
	//fmt.Println("This will not run")
}

func main() {
	safeDivision(10, 2)
	safeDivision(10, 0)
	fmt.Println("Program continues after panic")
	safeDivision(20, 2)
	//example:2
	demo()
}
