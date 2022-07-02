package main

import "fmt"

func main() {
	fmt.Println("slice example")

	var numSlice []int = []int{1, 2, 3}
	var slicePtr *[]int
	slicePtr = &numSlice
	fmt.Println(numSlice)

	var stringSlice = make([]string, 0)
	var strSlicePtr *[]string = &stringSlice
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, "1", "2", "3")
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, "5", "6", "7")
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice))
	stringSlice = append(stringSlice, "5", "6", "7")
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice))

	fmt.Println(len(*slicePtr), cap(*slicePtr))
	fmt.Println(len(*strSlicePtr), cap(*strSlicePtr))
}
