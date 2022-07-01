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

	// Type checker
	CheckType(10)
	CheckType(10.5)
	CheckType(true)
	CheckType("hello")
	CheckType(fmt.Println)
}

func CheckType(val interface{}) {
	fmt.Printf("%T -> %v \n", val, val)

	switch val.(type) {
	case bool:
		fmt.Println("Val is bool")
	case int, int8, int16, int32, int64:
		fmt.Println("Val is a number")
	case float32, float64:
		fmt.Println("Val is float")
	case string:
		fmt.Println("Val is string")
	default:
		fmt.Println("Val type is ", fmt.Sprintf("%T", val))
	}
}
