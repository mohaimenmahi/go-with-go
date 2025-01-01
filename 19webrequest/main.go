package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com"

func main() {
	fmt.Println("Golang Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of Type: %T\n", response)

	defer response.Body.Close() // Close the response body, to avoid memory leaks. It is caller's responsibility to close the response body.

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
