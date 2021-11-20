// PROBLEM: Find square root of a number
// https://leetcode.com/problems/sqrtx/
// https://www.interviewbit.com/problems/square-root-of-integer/

package main

import "fmt"

// SOLUTION
func sqrt(int x) int {
	r := x
	for r*r > x {
		r = (r + x/r)/2
	}
	return r
}


func main() {
	// INPUT :
	x := 8

	// OUTPUT :
	fmt.Println(sqrt(x))
}