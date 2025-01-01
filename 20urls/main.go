package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com/mohaimenmahi?tab=repositories"

func main() {
	fmt.Println("Urls in Golang")

	fmt.Println("My Github URL is: ", myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Query: ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of Query Params: %T\n", qparams)
	fmt.Println("Params: ", qparams)
	fmt.Println("Param data: ", qparams["tab"])

	for _, val := range qparams {
		fmt.Println("Value: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "/mohaimenmahi",
		RawQuery: "tab=repositories",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("Another URL: ", anotherUrl)
}

/*
Output:
Urls in Golang
My Github URL is:  https://github.com/mohaimenmahi?tab=repositories
Scheme:  https
Host:  github.com
Path:  /mohaimenmahi
Query:  tab=repositories
Type of Query Params: url.Values
Params:  map[tab:[repositories]]
Param data:  [repositories]
Value:  [repositories]
Another URL:  https://github.com/mohaimenmahi?tab=repositories

*/
