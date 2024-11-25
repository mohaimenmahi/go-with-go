package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time in Golang")

	presentTime := time.Now()

	fmt.Println("The current time is: ", presentTime)

	fmt.Println(presentTime.Format("01/02/2006 03:04:05 Monday"))
	// this date time and days are fixed, you can't change it
	// 01/02/2006 => January 2nd, 2006

	createdDate := time.Date(2020, time.June, 10, 23, 0, 0, 0, time.UTC)

	fmt.Println("Created Date: ", createdDate.Format("02/01/2006 03:04:05 Monday"))
}
