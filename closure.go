package main

import "fmt"

func seqGen(seed int) func() int {
	inner := func() int {
		seed++
		return seed
	}
	return inner
}

func main() {
	func1 := seqGen(1)
	fmt.Println(func1(), func1(), func1(), func1())

	func2 := seqGen(10)
	fmt.Println(func2(), func2(), func2(), func2(), func2())
}
