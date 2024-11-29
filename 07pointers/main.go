package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers in Golang")

	// var ptr *int
	// fmt.Println("Value of the pointer is: ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of the pointer is: ", ptr)         // this will print the address in memory
	fmt.Println("Value of the actual pointer is: ", *ptr) // this will print the actual value stored at that address

	*ptr = *ptr * 2

	fmt.Println("New value of the actual pointer is: ", myNumber) // this will print the actual value stored at that address
}
