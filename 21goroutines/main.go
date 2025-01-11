package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // wait group is used to wait for the go routines to finish
// this is generally pointer

func main() {
	// greeter("Hello")
	// greeter("World")

	// the above code will run in sequence

	// lets run them using goroutines

	// go greeter("Hello")
	// greeter("World")

	// the above code: hello will not be printed but the worlds will be printed
	// by adding the go keyword, we just told the go runtime that:
	// "Just fire up the thread to execute the function with Hello"
	// but we have not told where to come back and never waited for it to come back

	// go greeter("Hello")
	// greeter("World")

	websiteLists := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}

	for _, website := range websiteLists {
		go getStatusCode(website)

		wg.Add(1) // this keep on adding go rountines number to the wait group
		// as there is only one call happening here
	}

	wg.Wait() // wait for all the go routines to finish, it is called to the end of the main function to wait for all the go routines to finish
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond) // sleep for 3 milliseconds, then print
// 		fmt.Println(s, i+1)
// 	}
// }

func getStatusCode(endoint string) {
	req, err := http.Get(endoint)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Printf("%d status code for website %s\n", req.StatusCode, endoint)

	wg.Done() // this will be called when the go routine is finished, need to add it manually
}
