// PROBLEM: Next permutation
// https://leetcode.com/problems/next-permutation/
// https://www.interviewbit.com/problems/next-permutation/

package main

import (
	"fmt"
)

// SOLUTION
func next_permutation(nums []int) {
	n := len(nums)
	var k, l int

	for k = n - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	if k < 0 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	} else {
		for l = n - 1; l > k; l-- {
			if nums[l] > nums[k] {
				break
			}
		}

		nums[k], nums[l] = nums[l], nums[k]

		for i := 1; k+i <= n-i; i++ {
			nums[i+k], nums[n-i] = nums[n-i], nums[i+k]
		}
	}

}

func main() {
	// INPUT :
	nums := []int{1, 1, 5}

	// OUTPUT :
	next_permutation(nums)
	fmt.Println(nums)
}
