// Topic: Pointers

/*
x := 10
Declares an integer variable x and assigns it the value 10.
p := &x
The & operator gives the memory address of x.
p is now a pointer to x, meaning p stores the address where x is located in memory.
fmt.Println("x:", x, "pointer to x:", p, "value at pointer:", *p)
Prints the value of x.
Prints the pointer p (the memory address of x).
Prints the value at the pointer p using the * (dereference) operator, which accesses the value stored at the address p points to (which is 10).
*/
package main

import "fmt"

func main() {
	/*
		Pointers are variables that store memory addresses.
		& operator gets the memory address of a variable.
		* operator dereferences a pointer to get the value it points to.
	*/
	const y = 20
	x := 10
	p := &x
	fmt.Println("x:", x, "pointer to x:", p, "value at pointer:", *p)
	fmt.Println("Value of p:", *p)
	fmt.Println("Value of x:", &x)
	*p = y
	fmt.Println("Value of x:", x)
	fmt.Println("Value of x:", &x)
}
