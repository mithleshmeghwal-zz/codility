package main

import (
	"fmt"
)

func main() {
	A := [][]int{
		{1, 3, 4},
		{2, 3, 1, 5},
		{2},
		{1},
		{},
	}
	for i := 0; i < len(A); i++ {
		fmt.Println(Solution(A[i]))
	}
}

// sum of n elements - sum of given array
// key is array has sequenced elements only missing 1
func Solution(A []int) int {
	n := len(A)
	if n == 0 {
		return 1
	}
	if n == 1 {
		if A[0] == 1 {
			return 2
		}
		if A[0] == 2 {
			return 1
		}
	}
	n1 := n + 1 //  for 1----n

	sumOfNElements := (n1 * (n1 + 1)) / 2
	sumOfArray := 0
	for i := range A {
		sumOfArray += A[i]
	}
	return sumOfNElements - sumOfArray
}
