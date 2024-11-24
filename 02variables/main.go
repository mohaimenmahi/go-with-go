package main

import "fmt"

const LoginToken string = "iyfie4556wcfwv"

// Capital L has a significance in Go
// It means that the variable is exported
// It can be accessed outside the package
// It is a public variable

func main() {
	var username string = "John Doe"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float64 = 255.455445112544845656
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// default values and some aliases as well
	var anotherVariable int
	fmt.Println(anotherVariable) // it will print 0, as nothing is assigned
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	var defaultString string
	fmt.Println(defaultString) // it will print nothing, as nothing is assigned
	fmt.Printf("Variable is of type %T \n", defaultString)

	var defaultBool bool
	fmt.Println(defaultBool) // it will print false, as nothing is assigned
	fmt.Printf("Variable is of type %T \n", defaultBool)

	// implicit type
	// lexer will take care of it and assign the type
	var website = "youtube.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T \n", website)

	// no var style
	// this is only allowed inside a function
	// outside a function, it will throw an error
	numberOfUsers := 3000000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}
