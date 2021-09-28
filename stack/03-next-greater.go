// PROBLEM: Next greater element
// https://leetcode.com/problems/next-greater-element-ii/
// https://practice.geeksforgeeks.org/problems/next-larger-element/0

package main

import "fmt"

// SOLUTION
func next_greater(nums []int) []int {
	n := len(nums)
	var stack []int
	result := make([]int, n)
	for i:=0; i<n; i++ {
		result[i] = -1
	}

	for i:=0; i<n*2; i++ {
		for len(stack)!=0 && (nums[stack[len(stack)-1]] < nums[i%n]) {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return result
}


func main() {
	// INPUT :
	nums := []int{1,2,3,4,3}

	// OUTPUT :
	fmt.Println(next_greater(nums))
}