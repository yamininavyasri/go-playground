/*
	different types of functions

| Function Type    | Description                             |
| ---------------- | --------------------------------------- |
| Basic            | Named, no params or returns             |
| With Parameters  | Takes arguments                         |
| Multiple Returns | Returns several values                  |
| Named Returns    | Named return variables                  |
| Variadic         | Takes variable number of args           |
| Anonymous        | Inline function without name            |
| IIFE             | Anonymous function executed immediately |
| Method           | Function bound to a type                |
| Pointer Receiver | Method that modifies the type           |
| Higher-Order     | Takes or returns a function             |
*/
package main

import "fmt"

// 1️.Basic Function
func greet() {
	fmt.Println("Hello, Go!")
}

// 2️. Function with Parameters and Return Value
func add(a, b int) int {
	return a + b
}

// 3️. Function with Multiple Return Values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 4️. Named Return Values
func multiplyAndAdd(a, b, c int) (mul int, total int) {
	mul = a * b * c
	total = a + b + c
	return // returns mul, total automatically
}

// 5️. Variadic Function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 6️. Higher-Order Function (takes another function as argument)
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 7️. Function Returning a Function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 8️. Methods (value receiver)
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// 9️. Pointer Receiver Method
func (c *Circle) Scale(factor float64) {
	c.radius = c.radius * factor
}

func main() {
	fmt.Println("----- 1. Basic Function -----")
	greet()

	fmt.Println("\n----- 2. Function with Parameters -----")
	fmt.Println("3 + 4 =", add(3, 4))

	fmt.Println("\n----- 3. Multiple Returns -----")
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	fmt.Println("\n----- 4. Named Return Values -----")
	mul, total := multiplyAndAdd(2, 3, 4)
	fmt.Println("Multiply:", mul, "Total:", total)

	fmt.Println("\n----- 5. Variadic Function -----")
	fmt.Println("Sum of all numbers:", sum(1, 2, 3, 4, 5))

	fmt.Println("\n----- 6. Anonymous Function -----")
	sub := func(a, b int) int { return a - b }
	fmt.Println("10 - 3 =", sub(10, 3))

	fmt.Println("\n----- 7. Higher-Order Function -----")
	addFunc := func(x, y int) int { return x + y }
	fmt.Println("operate(5, 7, addFunc) =", operate(5, 7, addFunc))

	fmt.Println("\n----- 8. Function Returning a Function -----")
	double := multiplier(2)
	fmt.Println("Double of 6 =", double(6))

	fmt.Println("\n----- 9. Methods and Pointer Receivers -----")
	c := Circle{radius: 5}
	fmt.Println("Area before scaling:", c.Area())
	c.Scale(2)
	fmt.Println("Area after scaling:", c.Area())

	fmt.Println("\n----- 10. Immediately Invoked Function (IIFE) -----")
	result := func(x int) int {
		return x * x
	}(5)
	fmt.Println("Square of 5 =", result)

	fmt.Println("\n----- 11. Defer Example -----")
	defer fmt.Println("Deferred: runs after main() completes")
	fmt.Println("Main function is executing...")
}
