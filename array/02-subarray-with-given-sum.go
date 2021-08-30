// PROBLEM: Subarray with given sum
// https://practice.geeksforgeeks.org/problems/subarray-with-given-sum/0

package main

import "fmt"

// SOLUTION
func subarray_sum(A []int, S int, N int) []int {
	sum := A[0]
	start := 0
	var end int

	for end = 1; end <= N; end++ {
		for sum > S && start < end-1 {
			sum = sum - A[start]
			start++
		}

		if sum == S {
			return []int{start + 1, end}
		}

		if end < N {
			sum = sum + A[end]
		}
	}

	return []int{-1}
}

func main() {
	// INPUT :
	A := []int{1, 2, 3, 7, 5}
	S := 12
	N := len(A)

	// OUTPUT :
	ss := subarray_sum(A, S, N)
	for _, p := range ss {
		fmt.Print(p, " ")
	}
	fmt.Println()
}
