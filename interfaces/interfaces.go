/*
Classification: Abstract type (behavioral type).

Definition: An interface defines a set of method signatures without implementing them.

Purpose: To describe behavior that types can implement.

Abstract Type: You cannot create an interface instance directly; you need a type that implements its methods.

Polymorphism: Interfaces enable dynamic dispatch, allowing different types to be used interchangeably if they satisfy the interface.
*/
package main

import "fmt"

/*
Combined Interface Examples:
- Basic interface
- Slice of interfaces
- Type assertion / Type switch
- Polymorphism
*/

// 1️. Interface Definition
type Shape interface {
	Area() float64
}

// 2️. Structs implementing Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	fmt.Println("=== Basic Interface ===")
	var s Shape = Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println()

	fmt.Println("=== Slice of Interfaces ===")
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		Triangle{Base: 4, Height: 5},
	}

	for _, shape := range shapes {
		fmt.Println("Shape Area:", shape.Area())
	}
	fmt.Println()

	fmt.Println("=== Type Assertion / Type Switch ===")
	for _, shape := range shapes {
		switch sh := shape.(type) {
		case Circle:
			fmt.Println("Circle radius:", sh.Radius)
		case Rectangle:
			fmt.Println("Rectangle width & height:", sh.Width, sh.Height)
		case Triangle:
			fmt.Println("Triangle base & height:", sh.Base, sh.Height)
		}
		fmt.Println("Area:", shape.Area())
	}
	fmt.Println()

	fmt.Println("=== Polymorphism Demo ===")
	for _, shape := range shapes {
		fmt.Printf("%T Area: %.2f\n", shape, shape.Area())
	}
}
