package main

import "fmt"

func fact(x int) int {
	if x < 2 {
		fmt.Printf("%d \n", x)
		return 1
	} else {
		fmt.Printf("%d * ", x)
		return x * fact(x-1)
	}
}

func fib(x int) int {
	if x < 2 {
		fmt.Printf("%d \n", x)
		return x
	}
	fmt.Printf("%d ? ", x)
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println("Enter a number")
	var num int
	fmt.Scanf("%d", &num)
	fmt.Println("Factorial of", num, " -> ", fact(num))
	fmt.Println("Fibonacci of", num, " -> ", fib(num))

}
