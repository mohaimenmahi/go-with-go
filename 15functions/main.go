package main

import "fmt"

func main() {
	// we can not define a function inside a function
	fmt.Println("Functions in golang")

	greeter()

	result := adder(3, 5)
	fmt.Println("result is: ", result)

	proResult := proAdder(1, 2, 3, 4, 5) // just like *args in python, gets comma separated values and turns into a slice by ...int type
	fmt.Println("proResult is: ", proResult)

	proResult, message := proAdderWithMessage(1, 2, 3, 4, 5)
	fmt.Println("proResult is: ", proResult, " and message is: ", message)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int { // values is now a slice of int
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func proAdderWithMessage(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "This is the total"
}

func greeter() {
	fmt.Println("This is from greeter function")
}
