/*
In Go, struct tags are metadata attached to struct fields.
When working with JSON, tags tell the encoding/json package how to encode or decode struct fields between Go and JSON.
Tag syntax
`json:"fieldName"`
json → the name of the tag (used by the JSON package).
fieldName → the name that appears in the JSON.
You can rename JSON fields:
FullName string `json:"full_name"`
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Navya", Age: 25}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data)) // {"name":"Navya","age":25}
	json.Unmarshal([]byte(data), &p)
	fmt.Println(p)
}
