// PROBLEM: Trapping rain water
// https://leetcode.com/problems/trapping-rain-water/
// https://www.interviewbit.com/problems/rain-water-trapped/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func trap(height []int) int {
	h1, h2 := 0, 0
	total := 0

	for l, r := 0, len(height)-1; l<r; {
		if height[l]<height[r] {
			h1 = int(math.Max(float64(h1), float64(height[l])))
			total += h1 - height[l]
			l++
		} else {
			h2 = int(math.Max(float64(h2), float64(height[r])))
			total += h2 - height[r]
			r--
		}
	}

	return total
}


func main() {
	// INPUT :
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	// OUTPUT :
	fmt.Println(trap(height))
}