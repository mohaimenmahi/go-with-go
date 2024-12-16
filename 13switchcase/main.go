package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch Case in Golang")

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
	case 4:
		fmt.Println("You can move 4 steps")
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and roll the dice again")
	default:
		fmt.Println("What was that!")
	}
}
