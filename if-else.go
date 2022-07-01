package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a num")
	if _, err := fmt.Scanf("%d", &num); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("num", num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FIZZBUZZ")
	} else if num%3 == 0 {
		fmt.Println("FIZZ")
	} else if num%5 == 0 {
		fmt.Println("BUZZ")
	} else {
		fmt.Println(num)
	}
}
