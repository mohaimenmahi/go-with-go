package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// the first method is responsible for receiving the value
	// <-chan means that the channel is only for receiving the value
	// this is receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// we are not allowed to close the channel from the receiver side
		// close(ch) // this will throw an error
		val, isChannelOpen := <-ch
		// fmt.Println("Data from channel:", <-myCh)

		fmt.Println("Data from channel:", val, isChannelOpen)

		// we can differentiate between the zero value and the channel close value by using the second return boolean value of the channel
		// which is isChannelOpen
		wg.Done()
	}(myCh, wg)

	// the second method is responsible for putting value into the channel
	// chan<- means that the channel is only for sending the value
	// this is send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		// the listener has only one line, but as we have added buffer value as 2 in the channel, it will not block the execution

		close(ch) // it will close the channel, but we can still read the value from the channel, but we cannot write more values into the channel

		// myCh <- 7 // this will throw an error as the channel is closed
		wg.Done()
	}(myCh, wg)

	// if we close the channel by not sending any value, then the listener will get the zero value of the channel type
	// but also, if we send 0 normally as a value without or before closing the channel, then the listener will get the 0 value
	// so, how can we differentiate between the assigned zero value and the zero value because of the channel close?

	// the answer is, we can use the second return value of the channel, which is a boolean value
	// if it is true, then the channel is open, if it is false, then the channel is closed
	wg.Wait()
}
