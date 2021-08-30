// PROBLEM: Missing number in array
// https://leetcode.com/problems/missing-number/
// https://practice.geeksforgeeks.org/problems/missing-number-in-array/0

package main

import "fmt"

// SOLUTION
func missing_number(nums []int) int {
	result := len(nums)
	i := 0

	for _, num := range nums {
		result ^= num ^ i
		i++
	}

	return result
}

func main() {
	// INPUT :
	nums := []int{3, 2, 0, 1}

	// OUTPUT :
	fmt.Println(missing_number(nums))
}
