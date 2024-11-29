package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")

	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Fruit List type is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Pineapple")
	fmt.Println("Fruit List: ", fruitList)

	fruitList = append(fruitList[1:3]) // [1:3] means from index 1 to 2, the last index is non-inclusive
	// the colon syntax is used to slice up the slice, means
	// if you want to make seperate part of the slice
	// that's when we use the colon syntax

	// fruitList = append(fruitList[:3]) // [:3] means from start to index 2
	// fruitList = append(fruitList[1:]) // [1:] means from index 1 to end

	fmt.Println("Fruit List: ", fruitList)

	// diffence between make and new??
	// make is used to create slices, maps, and channels
	// new is used to create pointers of those types

	highScores := make([]int, 4) // defined a slice of int with length 4, alloacted memory for 4 elements

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// we can not insert element at index 4, because the length of slice is 4

	// but we can append the element
	highScores = append(highScores, 555, 666, 777)
	// append basically reallocates the memory and copies the elements to the new memory location

	fmt.Println("High Scores: ", highScores)

	// sort

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores: ", highScores)

	// check if sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))

	// removing a value from slice based on index
	var courses = []string{"react", "angular", "ruby", "python", "swift", "java"}

	fmt.Println("Courses: ", courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	// append the slice from start to index and then append the slice from index+1 to end
	// basically we are removing the element at index 2

	fmt.Println("Courses after removing: ", courses)
}
