package main

import (
	"fmt"
)

func main() {
	A := [][]int{
		{1, 2},
		{3, 1, 2, 4, 3},
	}
	for i := 0; i < len(A); i++ {
		fmt.Println(Solution(A[i]))
	}
}

func Solution(A []int) int {
	n := len(A)
	if n == 2 {
		return Abs(A[0] - A[1])
	}
	var sumRight int
	for i := 1; i < n; i++ {
		sumRight = sumRight + A[i]
	}
	sumLeft := A[0]
	minimumDifference := Abs(Abs(sumRight) - Abs(sumLeft))

	for i := 1; i < n; i++ {
		if Abs(sumLeft-sumRight) < minimumDifference {
			minimumDifference = Abs(sumLeft - sumRight)
		}
		sumLeft += A[i]
		sumRight -= A[i]
	}
	return minimumDifference
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
