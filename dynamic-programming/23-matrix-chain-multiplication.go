// PROBLEM: Burst Baloons (Matrix Chain Multiplication)
// https://leetcode.com/problems/burst-balloons/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func max_coins(nums []int) int {
    N := make([]int, len(nums)+2)
    n := 1
    for _, x := range nums {if x>0 {N[n]=x; n++}}
    N[0], N[n] = 1, 1
    n++

    dp := make([][]int, n)
    for i := range dp {dp[i]=make([]int, n)}

    for k:=2; k<n; k++ {
        for l:=0; l<n-k; l++ {
            r := l+k
            for i:=l+1; i<r; i++ {
                dp[l][r] = int(math.Max(float64(dp[l][r]), float64(N[l] * N[i] * N[r] + dp[l][i] + dp[i][r])))
            }
        }
    }

    return dp[0][n-1]
}


func main() {
    // INPUT :
    nums := []int{3,1,5,8}

    // OUTPUT :
    fmt.Println(max_coins(nums))
}