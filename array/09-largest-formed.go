// PROBLEM: Largest number formed from an array
// https://leetcode.com/problems/largest-number/
// https://www.interviewbit.com/problems/largest-number/
// https://practice.geeksforgeeks.org/problems/largest-number-formed-from-an-array/0

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// SOLUTION
func largest_formed(nums []int) string {
	var s []string
	for _, i := range nums {
		s = append(s, strconv.Itoa(i))
	}

	sort.Slice(s, func(i, j int) bool {
		a, b := s[i], s[j]
		return a+b > b+a
	})

	if s[0] == "0" {
		return "0"
	}

	result := strings.Builder{}
	for _, i := range s {
		result.WriteString(i)
	}

	return result.String()
}

func main() {
	// INPUT :
	nums := []int{3, 30, 34, 5, 9}

	// OUTPUT :
	fmt.Println(largest_formed(nums))
}
