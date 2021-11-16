// PROBLEM: Serach in a rotated sorted array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// https://www.interviewbit.com/problems/rotated-sorted-array-search/
// https://practice.geeksforgeeks.org/problems/search-in-a-rotated-array/0

package main

import "fmt"

// SOLUTION
func search_rotated(nums []int, target int) int {
    n := len(nums)
    l, h := 0, n-1

    for l<h {
        m := (l+h)/2
        if (nums[m]>nums[h]) {
            l = m+1
        } else { h = m }
    }

    rot := l
    l=0
    h=n-1
    for l<=h {
        m := (l+h)/2;
        rm := (m+rot)%n;
        if nums[rm]==target {return rm}
        if nums[rm]<target {
            l = m+1
        } else { h = m-1 }
    }

    return -1;
}


func main() {
    // INPUT :
    nums := []int{4,5,6,7,0,1,2}
    target := 0

    // OUTPUT :
    fmt.Println(search_rotated(nums, target))
}