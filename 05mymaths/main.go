package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths in Golang")

	// var myNumber int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("Addition: ", myNumber+int(myNumberTwo))

	// random number

	//fmt.Println(rand.Intn(5)) // it gives 0 to 5, this is like [0, 5)]

	// for getting 1 to 6
	//fmt.Println(rand.Intn(5) + 1) // it gives 1 to 6, this is like [1, 6)]

	// rand from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println("My randomNumber: ", myRandomNum)
}
