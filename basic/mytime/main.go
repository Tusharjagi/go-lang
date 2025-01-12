package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study of go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// important format syntax
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.July, 10, 23, 43, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}