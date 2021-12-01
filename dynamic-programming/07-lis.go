// PROBLEM: Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
// https://www.interviewbit.com/problems/longest-increasing-subsequence/

package main

import "fmt"

// SOLUTION
func lis(nums []int) int {
    tails := make([]int, len(nums))
    size := 0

    for _, x := range nums {
        i, j := 0, size

        for i!=j {
            m := int((i+j)/2)
            if tails[m]<x {
                i = m+1
            } else {
                j = m
            }
        }

        tails[i] = x
        if i==size {size++}
    }

    return size
}


func main() {
    // INPUT :
    nums := []int{10,9,2,5,3,7,101,18}

    // OUTPUT :
    fmt.Println(lis(nums))
}