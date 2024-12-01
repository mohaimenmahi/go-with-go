package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in Golang")

	languages := make(map[string]string) // convention is map[key_type]value_type

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of Languages:", languages)

	fmt.Println("JS Short for:", languages["JS"])

	// deleting value from map, same thing will be for slices as well
	delete(languages, "RB") // delete(map_name, key)

	fmt.Println("List of Languages after deleting Ruby:", languages)

	// loops in maps

	for key, value := range languages {
		fmt.Printf("For key %v the value is %v\n", key, value) // %v is used for any type of value
	}
}
