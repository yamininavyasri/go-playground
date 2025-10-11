package main

import (
	"fmt"

	"github.com/yaminivenkat.grandhi/go-playground/packages/mathutil"
	"github.com/yaminivenkat.grandhi/go-playground/packages/stringutil"
)

func main() {
	sum := mathutil.Add(5, 3)
	product := mathutil.Multiply(5, 3)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)

	reversed := stringutil.Reverse("Hello")
	upper := stringutil.Upper("world")
	fmt.Println("Reversed:", reversed)
	fmt.Println("Uppercase:", upper)
}
