// PROBLEM: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func mps(nums []int) int {
    n, l, r := len(nums), 0, 0
    result := nums[0]

    for i:=0; i<n; i++ {
        l = func()int{if l!=0 {return l} else {return 1}}() * nums[i]
        r = func()int{if r!=0 {return r} else {return 1}}() * nums[n-1-i]
        result = int(math.Max(float64(result), math.Max(float64(l), float64(r))))
    }

    return result
}


func main() {
    // INPUT :
    nums := []int{2,3,-2,4}

    // OUTPUT :
    fmt.Println(mps(nums))
}