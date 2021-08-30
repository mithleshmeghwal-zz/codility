package main

import (
	"fmt"
)

func main() {
	A := [][]int{
		{1, 3, 1, 4, 2, 3, 5, 4},
	}
	X := []int{
		5,
	}
	for i := 0; i < len(A); i++ {
		fmt.Println(Solution(X[i], A[i]))
	}
}
// use map and steps
func Solution(X int, A []int) int {
	steps := X
	m := map[int]bool{}
	for i := range A {
		if _, ok := m[A[i]]; !ok {
			m[A[i]] = true
			steps--
			if steps == 0 {
				return i
			}
		}
	}
	return -1
}
