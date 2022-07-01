package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	if _, err := fmt.Scanf("%d", &i); err != nil {
		fmt.Println("Error ", err)
		return
	}
	switch {
	case i%3 == 0 && i%5 == 0:
		fmt.Println("FIZZBUZZ")
	case i%3 == 0:
		fmt.Println("FIZZ")
	case i%5 == 0:
		fmt.Println("BUZZ")
	default:
		fmt.Println(i)
	}
	fmt.Println(time.Now())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend time")
	default:
		fmt.Println("Weekday. please work")

	}
}
