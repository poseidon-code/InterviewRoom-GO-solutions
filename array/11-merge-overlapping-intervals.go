// PROBLEM: Merge overlapping intervals
// https://leetcode.com/problems/merge-intervals/
// https://www.interviewbit.com/problems/merge-intervals/

package main

import (
	"fmt"
	"math"
	"sort"
)

// SOLUTION
func merge_intervals(intervals [][]int) [][]int {
	if len(intervals)<=1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i][0], intervals[j][0]
		return a<=b
	})

	var result [][]int
	result = append(result, intervals[0])
	
	for i:=1; i<len(intervals); i++ {
		if result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else {
			result[len(result)-1][1] = int(math.Max(float64(result[len(result)-1][1]), float64(intervals[i][1])))
		}
	}

	return result
}


func main() {
	// INPUT :
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}

	// OUTPUT :
	fmt.Println(merge_intervals(intervals))
}