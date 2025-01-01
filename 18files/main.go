package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang!")

	content := "This needs to go in a file and it is a long content"

	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err) // panic will just shut down the program and will show the error
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Printf("All done with the file of %v characters \n", length)

	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	// whenever we read the file, it will not being read in string format
	// specially when we read data from the internet, it will be in bytes format

	checkNilErr(err)

	fmt.Println("Text data in the file is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
