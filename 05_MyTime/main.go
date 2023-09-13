package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// "01-02-2006" = Date
	// "01-02-2006 Monday" = Includes Day
	// "01-02-2006 15:04:05 Monday" = Includes Time

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
