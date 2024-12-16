package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for loop

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days { // i is index
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days { // i is index, day is value, this is kind of for each loop in golang
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }

	for _, day := range days { // _ is used to ignore the index
		fmt.Println(day)
	}

	rougueValue := 1

	for rougueValue < 10 { // same as while loop in the other languages
		if rougueValue == 2 {
			goto lco // this will jump to the lco label
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumped to LCO")
}
