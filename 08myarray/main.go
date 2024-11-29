package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List Length is: ", len(fruitList))

	// define and initialize an array same time
	var vegList = [3]string{"Potato", "Beans", "Onion"}

	fmt.Println("Veg List is: ", vegList)
}
