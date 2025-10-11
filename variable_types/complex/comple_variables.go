/*
| Complex Type      | Purpose / Use                                        |
| ----------------- | ---------------------------------------------------- |
| Array             | Fixed-size sequence of elements                      |
| Slice             | Dynamic-size, flexible sequence                      |
| Map               | Key-value store, efficient lookup                    |
| Pointer           | Reference to a variable for modification             |
| Struct            | Group fields into a single type                      |
| Method            | Function bound to a type                             |
| Function variable | Store function in a variable, higher-order functions |
| Slice of structs  | Collection of objects                                |
| Map of slices     | Categorized collection of slices                     |
| Interface         | Abstract behavior, polymorphism                      |
| Channel           | Concurrency-safe communication                       |
*/
package main

import "fmt"

// Struct
type Person struct {
	Name string
	Age  int
}

// struct 2
type Person2 struct {
	name  string
	place string
}

// Method on struct
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I am %s and I am %d years old", p.Name, p.Age)
}

/*
	Implement fmt.Stringer

fmt.Stringer is an interface in Go:

	type Stringer interface {
	    String() string
	}

Any type that implements a String() string method satisfies this interface.
This is used by fmt.Println or fmt.Sprintf to get a custom string representation of your type.
*/
func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func (p1 Person2) details() string {
	return fmt.Sprintf("Hello, I am %s and I am from %s", p1.name, p1.place)
}

func main() {
	// Array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []string{"apple", "banana", "cherry"}
	slice = append(slice, "date")
	fmt.Println("Slice:", slice)

	// Map
	m := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	m["Charlie"] = 35
	fmt.Println("Map:", m)

	// Pointer
	x := 10
	p := &x
	*p = 20
	fmt.Println("Pointer:", *p, "Original x:", x)

	// Struct
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Struct:", person)
	fmt.Println(person.Greet())

	person2 := Person2{name: "Bob", place: "New York"}
	fmt.Println("Struct2:", person2)
	fmt.Println(person2.details())

	// Function as variable
	multiply := func(a, b int) int { return a * b }
	fmt.Println("Function variable multiply(3,4):", multiply(3, 4))

	// Higher-order function
	operate := func(a, b int, f func(int, int) int) int {
		return f(a, b)
	}
	fmt.Println("Higher-order function operate:", operate(5, 6, multiply))

	// Slice of structs
	people := []Person{
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	fmt.Println("Slice of structs:", people)

	// Map of slices
	group := map[string][]Person{
		"team1": people,
	}
	fmt.Println("Map of slices:", group)

	// Interface
	var s fmt.Stringer = person

	// 5️⃣ Print using the Stringer interface
	fmt.Println("Using Stringer interface:", s)

	// 6️⃣ You can also print the struct directly
	// fmt.Println calls the String() method automatically
	fmt.Println("Printing struct directly:", person)
	// Channel
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println("Channel value:", <-ch)
}
