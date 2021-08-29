package main

import (
	"fmt"
)

func main() {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	fmt.Println(Solution(A, K))
}

// Cyclic Rotation
// copy the last element in temp variable temp = A[n-1]
// start from left end
// copy last element to second last
// second last to third last
// until second element
// Run this process K times to perform right rotation

func Solution(A []int, K int) []int {
	n := len(A)
	if n == 0 {
		return A
	}
	if n == K {
		return A
	}
	var temp int
	for k := 0; k < K; k++ {
		temp = A[n-1]
		for i := n - 1; i > 0; i-- {
			A[i] = A[i-1]
		}
		A[0] = temp
	}
	return A
}
