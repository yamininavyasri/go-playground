package main

import (
	"fmt"
)

func main() {
	// Boolean
	var isGoFun bool = true
	fmt.Println("Boolean:", isGoFun)

	// Integer types
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	fmt.Println("Integers:", i, i8, i16, i32, i64)

	// Unsigned integers
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	fmt.Println("Unsigned Integers:", u, u8, u16, u32, u64)

	// Floating point
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359
	fmt.Println("Floats:", f32, f64)

	// Complex numbers
	var c complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i
	fmt.Println("Complex Numbers:", c, c2)

	// String
	var s string = "Hello Go!"
	fmt.Println("String:", s)

	// Rune (alias for int32, represents Unicode code point)
	var r rune = 'A'
	fmt.Println("Rune:", r)

	// Byte (alias for uint8)
	var b byte = 255
	fmt.Println("Byte:", b)

	// Default zero-values
	var defaultInt int
	var defaultBool bool
	var defaultString string
	fmt.Println("Default values -> Int:", defaultInt, "Bool:", defaultBool, "String:", defaultString)
}
