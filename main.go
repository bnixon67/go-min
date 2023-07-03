package main

import "fmt"

func min(v int, va ...int) int {
	minValue := v

	fmt.Printf("min(%v", v)
	for _, n := range va {
		fmt.Printf(", %d", n)
		if n < minValue {
			minValue = n
		}
	}
	fmt.Printf(") = %v\n", minValue)

	return minValue
}

func main() {
	min(1)
	min(1, 0, 2)
	min(1, 2, 3)
	min(1, 2, 3, 4)
	min(4, 3, 2, 1)
	min(3, 1, 0, 1)
	min(1, 2, -1, 3)
}
