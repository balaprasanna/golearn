package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var total int64
	for _, v := range m {
		total += v
	}
	return total
}

func SumFloats(m map[string]float64) float64 {
	var total float64
	for _, v := range m {
		total += v
	}
	return total
}

// GenericSum Generic sum functions to handle int/float
func GenericSum[K comparable, V int64 | float64](m map[K]V) V {
	var total V
	for _, val := range m {
		total += val
	}
	return total
}

type Number interface {
	int | int64 | float64
}

func GenericAvg[K comparable, V Number](m map[K]V) V {
	var avg V
	for _, val := range m {
		avg += val
	}
	avg = avg / V(len(m))
	return avg
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Println("Without generics")
	fmt.Println(SumInts(ints))
	fmt.Println(SumFloats(floats))

	fmt.Println("With generics")
	fmt.Println(GenericSum[string, int64](ints))
	fmt.Println(GenericSum(floats))

	fmt.Println("With generics")
	fmt.Println(GenericAvg(ints))
	fmt.Println(GenericAvg(floats))
}
