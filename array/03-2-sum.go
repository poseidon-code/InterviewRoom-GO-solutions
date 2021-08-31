// PROBLEM: 2 sum
// https://leetcode.com/problems/two-sum/
// https://www.interviewbit.com/problems/2-sum/
// https://practice.geeksforgeeks.org/problems/key-pair/0

package main

import "fmt"

// SOLUTION
func two_sum(nums []int, target int) []int {
	tmap := make(map[int]int)

	for i, n := range nums {
		_, m := tmap[n]

		if m {
			return []int{tmap[n], i}
		} else {
			tmap[target-n] = i
		}
	}

	return nil
}

func main() {
	// INPUT
	nums := []int{2, 7, 11, 15}
	target := 9

	// OUTPUT :
	twosum := two_sum(nums, target)
	fmt.Printf("[%d, %d]\n", twosum[0], twosum[1])
}
