package main

import (
	"fmt"
	"testing"
)

func add(a int, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := add(5, 4)
	expected := 9
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	} else {
		fmt.Printf("passed Expected: %d, Actual: %d", expected, result)
	}
}
