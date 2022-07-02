package main

import "fmt"

func calc(a int, b int, ops string) int {
	fmt.Printf("Log: ops %v %v %v\n", ops, a, b)
	var ans int
	switch ops {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
	case "%":
		ans = a % b
	}
	return ans
}

func sum(numbers ...int) (total int) {
	fmt.Print(numbers)
	for _, val := range numbers {
		total += val
	}
	return total
}

func main() {
	fmt.Println(calc(1, 2, "+"))
	fmt.Println(calc(2, 2, "-"))
	fmt.Println(calc(4, 2, "*"))
	fmt.Println(calc(6, 2, "/"))
	fmt.Println(calc(8, 2, "%"))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum([]int{10, 20, 30, 40, 50}...))
}
