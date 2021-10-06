// PROBLEM: Sliding window maximum
// https://leetcode.com/problems/sliding-window-maximum/
// https://www.interviewbit.com/problems/sliding-window-maximum/

package main

import "fmt"

// SOLUTION
func maximum(nums []int, k int) []int {
	var buffer []int
	var result []int

	for i:=0; i<len(nums); i++ {
		for len(buffer)!=0 && nums[i]>=nums[buffer[len(buffer)-1]] {
			buffer = buffer[:len(buffer)-1]
		}
		buffer = append(buffer, i)

		if i>=k-1 {
			result = append(result, nums[buffer[0]])
		}
		if buffer[0]<=i-k+1 {
			buffer = buffer[1:]
		}
	}

	return result
}


func main() {
    // INPUT :
    nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3

    // OUTPUT :
    fmt.Println(maximum(nums, k))
}