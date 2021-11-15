// PROBLEM: Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package main

import "fmt"

// SOLUTION
func search_range(nums []int, target int) []int {
    i, j := 0, len(nums)-1
    result := []int{-1,-1}

    for i<j {
        mid := (i+j)/2
        if nums[mid]<target { 
            i = mid+1
        } else { j = mid }
    }
    if nums[i] != target {
        return result
    } else { result[0] = i }

    j = len(nums)-1
    for i<j {
        mid := (i+j)/2 + 1
        if nums[mid]>target { 
            j = mid-1
        } else { i = mid }
    }
    result[1] = j

    return result
}


func main() {
    // INPUT :
    nums := []int{5,7,7,8,8,10}
    target := 8

    // OUTPUT :
    fmt.Println(search_range(nums, target))
}