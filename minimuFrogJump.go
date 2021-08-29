package main

import (
	"fmt"
)

func main() {
	X := []int{10, 1} // starting point
	Y := []int{85, 1} // edning point
	D := []int{30, 3} // jump length
	for i := 0; i < len(X); i++ {
		fmt.Println(Solution(X[i], Y[i], D[i]))
	}
}

// minimum jumps required to reach Y

func Solution(X, Y, D int) int {
	if D > Y {
		return 0
	}
	Y = Y - X
	var add int
	if Y%D > 0 {
		add++
	}
	return (Y / D) + add
}
