/*
The defer statement delays the execution of a function until the surrounding function returns.
It’s often used for cleanup tasks (closing files, releasing resources, unlocking mutexes, etc.).
The deferred function runs after the current function completes.
Deferred functions execute in LIFO (Last In, First Out) order if multiple defer statements are used.
Commonly used for cleanup, like closing files, releasing locks, or printing log
basic syntax: defer functionName(arguments)
*/
package main

import "fmt"

func main() {
	fmt.Println("Start")

	//Notice: Last defer executes first.
	defer fmt.Println("This runs last")
	fmt.Println("Main function running")
	defer fmt.Println("First deferred")
	defer fmt.Println("Second deferred")
	defer fmt.Println("Third deferred")

	//Defer with arguments evaluated immediately
	x := 5
	defer fmt.Println("Deferred:", x)
	x = 10
	fmt.Println("x now:", x)
	fmt.Println("End")
}
