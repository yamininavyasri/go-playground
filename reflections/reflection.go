/*
Reflection is useful when:
You don’t know the type of data at compile time.
You’re writing generic functions or libraries (like JSON encoding, ORM frameworks, etc.).
You want to inspect struct fields, tags, or method names dynamically.
*/
package main

import (
	"fmt"
	"reflect"
)

// Define a struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create an instance of Person
	p := Person{Name: "Jack", Age: 25}

	// Get the reflect.Type and reflect.Value
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// Iterate over struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("\nField Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", field.Type)
		fmt.Printf("Field Value: %v\n", value)
		fmt.Printf("JSON Tag: %s\n", field.Tag.Get("json"))
	}

	// Check if a variable is of a certain kind
	fmt.Println("\nIs struct?", t.Kind() == reflect.Struct)
}
