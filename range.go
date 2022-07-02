package main

import "fmt"

func main() {
	var nums = []int{11, 22, 33}
	for _, v := range nums {
		fmt.Println(v)
	}
	nums = append(nums, 1, 2, 3)
	for _, v := range nums {
		fmt.Println(v)
	}
}
