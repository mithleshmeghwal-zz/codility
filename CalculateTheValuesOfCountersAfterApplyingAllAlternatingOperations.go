//Calculate the values of counters after applying all alternating operations:
//You are given N counters, initially set to 0, and you have two possible operations on them:

//increase(X) − counter X is increased by 1,
// max counter − all counters are set to the maximum value of any counter.

package main

import "fmt"

func main() {
	A := [][]int{
		{3, 4, 4, 6, 1, 4, 4},
	}
	N := []int{
		5,
	}
	for i := range A {
		fmt.Println(A[i], Solution(N[i], A[i]))
	}
}

func Solution(N int, A []int) []int {
	R := make([]int, N)
	max := 0
	min := 0
	for k := range A {
		if A[k] <= N {
			// counter X is increased by 1
			if R[A[k]-1] < min+1 {
				R[A[k]-1] = min + 1
			} else {
				R[A[k]-1]++
			}
			if R[A[k]-1] > max {
				max = R[A[k]-1]
			}
		} else {
			min = max
			// set all counters are set to the maximum value of any counter.

		}
		fmt.Println(A[k], R, max, min)
	}
	for i := range R {
		if R[i] < min {
			R[i] = min
		}
	}
	return R
}

