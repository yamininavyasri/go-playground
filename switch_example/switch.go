// Topic: Switch Statement in Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// basic switch
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Some other day")
	}

	// Switch with multiple expressions in a case
	color := "red"
	switch color {
	case "dark", "red":
		fmt.Println("Primary color", color)
	case "green":
		fmt.Println("Secondary color")
	default:
		fmt.Println("Unknown color")
	}

	// Switch without an expression (like if-else chain)
	num := -1
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0:
		fmt.Println("Positive number")
	}

	// Type switch
	/*
		var x interface{} = time.Now()
		interface{} is the empty interface in Go.
		It can hold a value of any type.
		time.Now() returns a value of type time.Time.
		So, x is a variable of type interface{} holding a time.Time value.
	*/
	var x interface{} = time.Now()
	x = 10
	switch v := x.(type) {
	case int:
		fmt.Println("x is int:", v)
	case string:
		fmt.Println("x is string:", v)
	case time.Time:
		fmt.Println("x is of type time.Time:", v)
	default:
		fmt.Println("Unknown type")
	}

}
