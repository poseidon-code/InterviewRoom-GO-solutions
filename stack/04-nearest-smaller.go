// PROBLEM: Nearest smaller
// https://www.interviewbit.com/problems/nearest-smaller-element/

package main

import "fmt"

// SOLUTION
func nearest_smaller(nums []int) []int {
	result := []int{-1}
	stack := []int{nums[0]}

	for i:=1; i<len(nums); i++ {
		for len(stack)!=0 && stack[len(stack)-1] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack)==0 {
			result = append(result, -1)
		} else {
			result = append(result, stack[len(stack)-1])
		}
		stack = append(stack, nums[i])
	}

	return result
}


func main() {
	// INPUT :
	nums := []int{4,5,2,10,8}

	// OUTPUT :
	fmt.Println(nearest_smaller(nums))
}