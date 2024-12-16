package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

func main() {
	fmt.Println("Methods in Golang")

	// There is not much of big difference with functions and methods in Golang
	// when we do regular things like passing arguments and returning values that is function
	// when the functions actually go into the classes, we call them methods
	// as there is no classes in Golang, we use structs to define methods

	mohaimen := User{"Mohaimen", "mmc@go.dev", true, 30}

	fmt.Println("User: ", mohaimen) // prints => User:  {Mohaimen mmc@go.dev true 30}

	fmt.Printf("mohaimen details are %+v\n", mohaimen)
	// prints => mohaimen details are {Name:Mohaimen Email:mmc@go.dev Status:true Age:30}

	fmt.Println("Name is", mohaimen.Name) // prints => Name is Mohaimen

	mohaimen.GetStatus() // prints => Is user active? true

	mohaimen.NewMail() // prints => Email of this user is:

	fmt.Println("Email of this user is:", mohaimen.Email) // prints => Email of this user is:
}
