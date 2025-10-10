package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for index, value := range [5]int{1, 2, 3, 4, 5} {
		fmt.Println(index, value)
	}
	fmt.Println("Slice value is : ", [5]int{1, 2, 3, 4, 5})

	// For loop as a while loop
	fmt.Println("\nFor loop as a while loop:")
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// Infinite loop (with break)
	fmt.Println("\nInfinite loop with break:")
	k := 0
	for {
		if k >= 2 {
			break
		}
		fmt.Println("k =", k)
		k++
	}

	// Looping over a slice with range
	fmt.Println("\nLooping over a slice with range:")
	nums := []int{10, 20, 30}
	fmt.Println("Slice value is : ", nums)
	for idx, val := range nums {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
	}

	// Looping over a map with range
	fmt.Println("\nLooping over a map with range:")
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("Map value is : ", m)
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Looping over a string with range

	str := "hello"
	fmt.Println("\nString value is : ", str)
	fmt.Println("Looping over a string with range:")
	for idx, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", idx, char)
	}

}
