// PROBLEM: Sort an array of 0s, 1s and 2s
// https://leetcode.com/problems/sort-colors/
// https://practice.geeksforgeeks.org/problems/sort-an-array-of-0s-1s-and-2s/0

package main

import "fmt"

// SOLUTION
func sort_array(nums []int) {
	l := 0
	m := 0
	h := len(nums) - 1

	for m <= h {
		if nums[m] == 0 {
			nums[m], nums[l] = nums[l], nums[m]
			m++
			l++
		} else if nums[m] == 1 {
			m++
		} else {
			nums[m], nums[h] = nums[h], nums[m]
			h--
		}
	}

	fmt.Println(nums)
}

func main() {
	// INPUT :
	nums := []int{2, 0, 2, 1, 1, 0}

	// OUTPUT :
	sort_array(nums)
}
