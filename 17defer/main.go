package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

/* Output of the main function:
Hello
43210Two
One
World
*/

// there is not defer statement before the myDefer function call, so it will be executed in the normal order
// evreything inside the myDefer function will be executed, even though there are defer statements, they will
// be executed by taking the myDefer function as reference function
// so the output of the myDefer function will be 43210
