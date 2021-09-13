package main

import "fmt"

func main() {
	A := [][]int{
		{4, 1, 3, 2},
		{4, 1, 3},
		{1, 1, 4}, // anti sum
	}
	for i := range A {
		fmt.Println(A[i], Solution(A[i]))
	}
}
// fund sum and compare with actual sum
// ** looks for duplicates
func Solution(A []int) int {
	visted := map[int]bool{}
	expectedN := 0
	for i := range A {
		if _, ok := visted[A[i]]; ok {
			return 0
		}
		visted[A[i]] = true
		expectedN = expectedN + A[i]
	}
	n := len(A)
	sumOfN := n * (n + 1) / 2
	if sumOfN == expectedN {
		return 1
	}
	return 0
}
