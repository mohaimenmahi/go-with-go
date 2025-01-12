package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // wait group is used to wait for the go routines to finish
// this is generally pointer

var mut sync.Mutex // mutex is used to lock the code, so that only one go routine can access the code at a time

var signals = []string{}

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
		"https://www.golang.org",
		"https://www.github.com",
	}

	for _, website := range websiteLists {
		go getStatusCode(website)

		wg.Add(1) // this keep on adding go rountines number to the wait group
		// as there is only one call happening here
	}

	wg.Wait() // wait for all the go routines to finish, it is called to the end of the main function to wait for all the go routines to finish
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond) // sleep for 3 milliseconds, then print
// 		fmt.Println(s, i+1)
// 	}
// }

func getStatusCode(endpoint string) {
	req, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error", err)
	}

	mut.Lock() // lock the memory, so that only one go routine can access the memory at a time
	signals = append(signals, endpoint)
	mut.Unlock() // unlock the memory, it is necessary to unlock the memory just after writing is done, otherwise it will be locked forever

	fmt.Printf("%d status code for website %s\n", req.StatusCode, endpoint)

	wg.Done() // this will be called when the go routine is finished, need to add it manually
}
