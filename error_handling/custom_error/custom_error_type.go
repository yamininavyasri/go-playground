/*
Custom type MyError implements the error interface by defining the Error() string method.
You can now return rich error information.
*/
package main

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func riskyOperation(flag bool) error {
	if flag {
		return &MyError{Code: 500, Message: "Something went wrong"}
	}
	return nil
}

func main() {
	err := riskyOperation(true)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded")
	}
}
