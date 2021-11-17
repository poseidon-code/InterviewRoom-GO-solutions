// PROBLEM: Find minimum in a rotated sorted array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

import "fmt"

// SOLUTION
func minimum(nums []int) int {
    start, end := 0, len(nums)-1
    
    for start<end {
        if nums[start]<nums[end] {
            return nums[start]
        }
        
        mid := (start+end)/2

        if nums[mid]>=nums[start] {
            start = mid+1
        } else { end = mid }
    }

    return nums[start]
}


func main() {
    // INPUT :
    nums := []int{3,4,5,1,2}

    // OUTPUT :
    fmt.Println(minimum(nums))
}