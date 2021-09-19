// given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
// {1, 3, 6, 4, 1, 2}, should return 5
// {1, 2, 3},          should return 4
// {-1, -3},           should return 1
package main

import "fmt"

func main() {
	A := [][]int{
		{1, 3, 6, 4, 1, 2},
		{1, 2, 3},
		{-1, -3},
	}
	for i := range A {
		fmt.Println(A[i], Solution(A[i]))
	}
}

func Solution(A []int) int {
	if len(A) == 1 {
		if A[0] < 1 {
			return 1
		}
		if A[0] == 1 {
			return 2
		}
		return 1 // A[0] > 1
	}
	m := make([]bool, len(A))
	for i := range A {
		if A[i] > 0 && A[i] < len(A) {
			m[A[i]] = true
		}
	}
	var smallest int
	for i := 1; i < len(m); i++ {
		if !m[i] {
			smallest = i
			break
		}
	}
	if smallest == 0 {
		return len(m) + 1
	}
	return smallest
}
