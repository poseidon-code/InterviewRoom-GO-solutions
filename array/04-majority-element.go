// PROBLEM: Majority element
// https://leetcode.com/problems/majority-element/
// https://www.interviewbit.com/problems/majority-element/
// https://practice.geeksforgeeks.org/problems/majority-element/0

package main

import "fmt"

// SOLUTION
func majority_element(nums []int) int {
	major := nums[0]
	count := 1

	for _, n := range nums {
		if count == 0 {
			count++
			major = n
		} else if major == n {
			count++
		} else {
			count--
		}
	}

	return major
}

func main() {
	// INPUT :
	nums := []int{3, 2, 3}

	// OUTPUT :
	fmt.Println(majority_element(nums))
}
