// PROBLEM: Maximum subarray
// https://leetcode.com/problems/maximum-subarray/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func max_subarray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    result := dp[0]

    for i:=1; i<n; i++ {
        dp[i] = nums[i] + func()int{if dp[i-1]>0 {return dp[i-1]} else {return 0}}()
        result = int(math.Max(float64(result), float64(dp[i])))
    }

    return result
}


func main() {
    // INPUT :
    nums := []int{5,4,-1,7,8}

    // OUTPUT :
    fmt.Println(max_subarray(nums))
}