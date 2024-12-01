package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Frequent User"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 { // sometimes we need to initialize a variable and use it in the if block
		fmt.Println("This number is less than 10")
	} else {
		fmt.Println("This number is not less than 10")
	}

}
