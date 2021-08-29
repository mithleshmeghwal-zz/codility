package main

import (
	"fmt"
)

func main() {
	A := 1041
	fmt.Println(Solution(A))
}

// BinaryGAP
// hint use right binary shift

func Solution(N int) int {
	start := false
	result := 0
	result_temp := 0
	for N > 0 {
		if N%2 == 1 {
			if start {
				if result_temp > result {
					result = result_temp
					result_temp = 0
				}
			} else {
				start = true
				result_temp = 0
			}
		} else {
			if start {
				result_temp++
			}
		}
		N = N / 2
	}
	return result
}
