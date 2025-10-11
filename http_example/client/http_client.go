/*
This Go program demonstrates making an HTTP GET request to an external API (https://api.github.com) and reading the response.
It covers basic HTTP client usage, error handling, and reading response data in Go.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	fmt.Println(string(body))
}
