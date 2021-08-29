package main

import (
	"fmt"
)

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(A))
}

// lopp array, use value as index store element in map,

func Solution(A []int) int {
	m := map[int]int{}
	for i := range A {
		if _, ok := m[A[i]]; ok {
			delete(m, A[i])
		} else {
			m[A[i]] = 1
		}
	}
	for i := range m {
		return i
	}
	return -1
}
