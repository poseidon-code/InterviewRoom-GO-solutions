// PROBLEM: Largest rectangle in a histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/
// https://www.interviewbit.com/problems/largest-rectangle-in-histogram/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func rectangle(heights []int) int {
	var stack []int
	l := len(heights)
	max_area := 0

	for i:=0; i<=l; i++ {
		var h int
		if i==l {h=0} else {h=heights[i]}

		if len(stack)==0 || h>=heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			tp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var t int
			if len(stack)==0 {t=i} else {t=i-1-stack[len(stack)-1]}
			max_area = int(math.Max(float64(max_area), float64(heights[tp] * t)))
			i--
		}
	}

	return max_area
}

func main() {
	// INPUT :
	heights := []int{2,1,5,6,2,3}

	// OUTPUT :
	fmt.Println(rectangle(heights))
}