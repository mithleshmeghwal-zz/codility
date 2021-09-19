// given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:
// { i : A ≤ i ≤ B, i mod K = 0 }
// For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.package main

package main

import "fmt"

func main() {
	A := []int{
		6, 0, 10,
	}
	B := []int{
		11, 1, 10,
	}
	K := []int{
		2, 11, 5,
	}
	for i := range A {
		fmt.Println(A[i], B[i], K[i], Solution(A[i], B[i], K[i]))
	}
}
func Solution(A int, B int, K int) int {
	count := 0
	if A == 0 || A%K == 0 {
		count++
	}
	return B/K - A/K + count
}
