// PROBLEM: Longest palindromic subsequence
// https://leetcode.com/problems/longest-palindromic-subsequence/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func lps(s string) int {
    n := len(s)
    dp := make([][]int, n+1)
    for i := range dp {dp[i] = make([]int, n)}
    for i:=0; i<n; i++ {dp[1][i]=1}

    for i:=2; i<=n; i++ {
        for j:=0; j<n-i+1; j++ {
            dp[i][j] = func()int{
                if s[j]==s[i+j-1] {
                    return 2+dp[i-2][j+1] 
                } else {
                    return int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j+1])))
                }
            }()
        }
    }

    return dp[n][0]
}


func main() {
    // INPUT :
    s := "bbbab"

    // OUTPUT :
    fmt.Println(lps(s))
}