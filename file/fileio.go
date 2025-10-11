/*
File I/O in Go
File I/O allows a Go program to read from and write to files. The standard library provides the os and io/ioutil (deprecated in Go 1.16) or io & bufio packages for this.
| Operation      | Function                                                                 |
| -------------- | ------------------------------------------------------------------------ |
| Open a file    | `os.Open("filename")`                                                    |
| Create a file  | `os.Create("filename")`                                                  |
| Write to file  | `file.Write([]byte("text"))`                                             |
| Read from file | `file.Read(buffer)` or `io.ReadAll(file)`                                |
| Close a file   | `file.Close()`                                                           |
| Buffered I/O   | Use `bufio.NewWriter(file)` or `bufio.NewReader(file)` for efficient I/O |
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// --- Writing to a file ---
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, Go File I/O!\n")
	writer.WriteString("Writing multiple lines is easy.\n")
	writer.Flush() // flush buffered data to file

	fmt.Println("Data written to example.txt")

	// --- Reading from a file ---
	readFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	fmt.Println("Reading file content:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// --- Deleting the file ---
	// uncomment below code to see the file created.
	// err = os.Remove("example.txt")
	// if err != nil {
	// 	fmt.Println("Error deleting file:", err)
	// 	return
	// }
	// fmt.Println("File deleted successfully:", "example.txt")
}
