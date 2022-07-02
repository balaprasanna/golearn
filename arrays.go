package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [10]int
	for j := 0; j < 10; j++ {
		array[j] = rand.Int()
		fmt.Println(rand.Int())
	}
	for _, i := range array {
		fmt.Println(i * i)
	}
}
