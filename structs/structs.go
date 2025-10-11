/*
Classification: Concrete, composite data type.

Definition: A struct is a composite type that groups together fields (variables) under a single name.

Purpose: To represent real-world entities or data objects.

Concrete Type: You can create instances of a struct directly.

Can have methods: You can attach methods to a struct (value or pointer receivers).
*/
package main

import "fmt"

/*
Combined Struct Examples:
- Basic struct
- Methods (value & pointer receivers)
- Embedding / Composition
- Embedding + Interface
*/

// 1️. Basic Struct
type Person struct {
	Name string
	Age  int
}

// 2️. Methods

// Value receiver: does not modify original
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I am %s", p.Name)
}

// Pointer receiver: modifies original struct
func (p *Person) HaveBirthday() {
	p.Age += 1
}

// 3️. Struct Embedding / Composition
type Employee struct {
	Person   // Embedded struct
	Position string
}

// 4️. Embedding + Interface
type Greeter interface {
	Greet() string
}

func main() {
	fmt.Println("=== Basic Struct ===")
	p := Person{Name: "Alice", Age: 25}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println()

	fmt.Println("=== Methods ===")
	fmt.Println(p.Greet())
	p.HaveBirthday()
	fmt.Println("Age after birthday:", p.Age)
	fmt.Println()

	fmt.Println("=== Struct Embedding / Composition ===")
	e := Employee{
		Person:   Person{Name: "Bob", Age: 30},
		Position: "Developer",
	}
	fmt.Println("Employee Name:", e.Name) // Access embedded field directly
	fmt.Println("Age:", e.Age)
	fmt.Println("Position:", e.Position)
	fmt.Println()

	fmt.Println("=== Embedding + Interface ===")
	var g Greeter = e.Person // Employee satisfies Greeter through embedded Person
	fmt.Println(g.Greet())
}
