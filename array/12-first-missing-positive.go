// PROBLEM: First missing positive
// https://leetcode.com/problems/first-missing-positive/
// https://www.interviewbit.com/problems/first-missing-integer/

package main

import "fmt"

// SOLUTION
func missing_positive(nums []int) int {
	n := len(nums)

	for i:=0; i<n; i++ {
		for nums[i]>0 && nums[i]<=n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i:=0; i<n; i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}

	return n+1
}


func main() {
	// INPUT :
	nums := []int{7,8,9,11,12,-3}

	// OUTPUT :
	fmt.Println(missing_positive(nums))
}