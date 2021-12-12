// PROBLEM: Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/

package main

import (
	"fmt"
)

// SOLUTION
func partition(nums []int) bool {
    sum := 0
    for _, n := range nums {sum += n}
    if (sum & 1) == 1 {return false}
    sum /= 2

    dp := make([]bool, sum+1)
    dp[0] = true

    for _, x := range nums {
        for i:=sum; i>0; i-- {
            if i >= x {
                dp[i] = dp[i] || dp[i-x]
            }
        }
    }

    return dp[sum]
}


func main() {
    // INPUT :
    nums := []int{1,5,11,5}

    // OUTPUT :
    fmt.Println(partition(nums))
}