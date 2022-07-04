package main

import "fmt"

type NumberType interface {
	int | int64 | float64
}

func GenericCalculator[T NumberType](a T, b T, ops string) T {
	fmt.Printf("Log: ops %v %v %v\n", ops, a, b)
	var ans T
	switch ops {
	case "+":
		ans = a + b
	case "-":
		ans = a - b
	case "*":
		ans = a * b
	case "/":
		ans = a / b
		//case "%":
		//	ans = a % b
	}
	return ans
}

func main() {
	fmt.Println(GenericCalculator(1, 2, "+"))
	fmt.Println(GenericCalculator(2, 2, "-"))
	fmt.Println(GenericCalculator(4, 2, "*"))
	fmt.Println(GenericCalculator(6, 2, "/"))
	//fmt.Println(GenericCalculator(8, 2, "%"))
}
