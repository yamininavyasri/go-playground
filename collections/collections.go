// collections in go : arrays, slices, maps
package main

import "fmt"

func main() {

	fmt.Println("Collections")
	//declare an array

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr, "length of array is : ", len(arr))

	//declare an array with values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2, "length of array is : ", len(arr2))

	//declare a slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, "length of slice1 is : ", len(slice))

	//declare a slice with make
	slice2 := make([]int, 5)
	fmt.Println(slice2, "length of slice2 is : ", len(slice2))

	//declare a slice with make and capacity, make([]T, length, capacity) -> T indicates the type of the slice, length indicates the length of the slice, capacity indicates the capacity of the slice
	//If the slice is full (length = capacity), Go will automatically allocate a new underlying array (usually doubling capacity).
	slice3 := make([]int, 5, 10)
	fmt.Println(slice3, "length of slice3 is : ", len(slice3))

	//declare a slice with append
	slice4 := append(slice, 6)
	fmt.Println(slice4, "length of slice4 is : ", len(slice4))

	//declare a slice with append and capacity
	slice5 := append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice5, "length of slice5 is : ", len(slice5))

	//declare a map
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m, "length of map1 is : ", len(m))

	//declare a map with make
	m2 := make(map[string]int)
	fmt.Println(m2, "length of map2 is : ", len(m2))

	//declare a map with make and capacity
	m3 := make(map[string]int, 5)
	fmt.Println(m3, "length of map3 is : ", len(m3))

	//declare a map with make and capacity and values
	m4 := make(map[string]int, 5)
	m4["a"] = 1
	m4["b"] = 2
	fmt.Println(m4, "length of map4 is : ", len(m4))
}
