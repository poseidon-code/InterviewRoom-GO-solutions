// PROBLEM: Length of longest valid parentheses
// https://leetcode.com/problems/longest-valid-parentheses/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func longest_length(s string) int {
	var parentheses []int
	parentheses = append(parentheses, -1)
	result := 0

	for i:=0; i<len(s); i++ {
		if s[i]=='(' {
			parentheses = append(parentheses, i)
		} else {
			parentheses = parentheses[:len(parentheses)-1]

			if len(parentheses)==0 {
				parentheses = append(parentheses, i)
			} else {
				result = int(math.Max(float64(result), float64(i - parentheses[len(parentheses)-1])))
			}
		}
	}

	return result
}


func main() {
	// INPUT :
	s := ")()())"

	// OUTPUT :
	fmt.Println(longest_length(s))
}