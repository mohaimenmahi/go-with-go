package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza: ")

	// comma ok || error ok syntax
	// in the golang, there is no try-catch
	// if something goes wrong, somebody needs to catch that
	// in the case of the error, we can use the comma ok syntax
	input, _ := reader.ReadString('\n') // '\n' is the delimiter, it will read until the user hits enter

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)
}
