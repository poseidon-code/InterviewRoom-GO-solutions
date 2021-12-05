// PROBLEM: House Robber
// https://leetcode.com/problems/house-robber/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func rob(nums []int) int {
    a, b, n := 0, 0, len(nums)
    
    for i:=0; i<n; i++ {
        if i%2==0 {
            a = int(math.Max(float64(a+nums[i]), float64(b)))
        } else {
            b = int(math.Max(float64(a), float64(b+nums[i])))
        }
    }

    return int(math.Max(float64(a), float64(b)))
}


func main() {
    // INPUT :
    nums := []int{1,2,3,1}

    // OUTPUT :
    fmt.Println(rob(nums))
}