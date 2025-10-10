// Topic: Conditions
package main

import "fmt"

func main() {
	// if statement - x is block scoped, cannot be accessible outside of if block.
	if x := 1; x > 5 {
		fmt.Println("x is greater than 5", x)
	} else {
		fmt.Println("x is 5 or less ", x)
	}
	fmt.Println("x is not accessible outside of if block")
	// fmt.Println(x)

	x := 10
	if x > 20 {
		fmt.Println("x is greater than 20")
	} else if x > 10 {
		fmt.Println("x is greater than 10 but less than or equal to 20")
	} else {
		fmt.Println("x is 10 or less")
	}

}
