// PROBLEM: Find duplicate number
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func duplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i])) - 1)

		nums[index] *= -1

		if nums[index] > 0 {
			return int(math.Abs(float64(nums[i])))
		}
	}

	return -1
}

func main() {
	// INPUT :
	nums := []int{3, 1, 3, 4, 2}

	// OUTPUT :
	fmt.Println(duplicate(nums))
}
