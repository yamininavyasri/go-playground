//go:build prod
// +build prod

package main

import "fmt"

func printMessage() {
	fmt.Println("Running in Production Mode")
}
