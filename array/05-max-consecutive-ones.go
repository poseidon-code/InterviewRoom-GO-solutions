// PROBLEM: Max consecutive ones
// https://leetcode.com/problems/max-consecutive-ones/
// https://www.interviewbit.com/problems/max-continuous-series-of-1s/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func max_ones(nums []int) int {
	var max float64 = 0
	var count float64 = 0

	for _, n := range nums {
		if n == 1 {
			count++
			max = math.Max(count, max)
		} else {
			count = 0
		}
	}

	return int(max)
}

func main() {
	// INPUT :
	nums := []int{1, 1, 0, 1, 1, 1}

	// OUTPUT :
	fmt.Println(max_ones(nums))
}
