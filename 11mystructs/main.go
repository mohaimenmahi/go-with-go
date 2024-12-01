package main

import "fmt"

// this is like a class in other languages
// uppercase is public, lowercase is private
// lowercase properties are private
// uppercase properties are public

// defining this anywehere in the code will make it available to all the files in the package
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Stucts in Golang")
	// there is no inheritance in Golang; No super or parent

	mohaimen := User{"Mohaimen", "mmc@go.dev", true, 30}

	fmt.Println("User: ", mohaimen) // prints => User:  {Mohaimen mmc@go.dev true 30}

	fmt.Printf("mohaimen details are %+v\n", mohaimen)
	// prints => mohaimen details are {Name:Mohaimen Email:mmc@go.dev Status:true Age:30}

	fmt.Println("Name is", mohaimen.Name) // prints => Name is Mohaimen
}
