// PROBLEM: Longest common substring
// https://leetcode.com/problems/maximum-length-of-repeated-subarray/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func lcss(nums1, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    dp := make([]int, n+1)
    result := 0

    for i:=1; i<=m; i++ {
        for j:=n; j>0; j-- {
            if nums1[i-1] == nums2[j-1] {
                dp[j] = 1 + dp[j-1]
                result = int(math.Max(float64(result), float64(dp[j])))
            } else {
                dp[j] = 0
            }
        }
    }

    return result
}


func main() {
    // INPUT :
    nums1 := []int{1,2,3,2,1}
    nums2 := []int{3,2,1,4,7}

    // OUTPUT :
    fmt.Println(lcss(nums1, nums2))
}