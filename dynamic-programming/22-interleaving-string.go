// PROBLEM: Interleaving String
// https://leetcode.com/problems/interleaving-string/
// https://www.interviewbit.com/problems/interleaving-strings/

package main

import (
	"fmt"
)

// SOLUTION
func is_interleaving(s1, s2, s3 string) bool {
    m, n := len(s1), len(s2)
    if len(s3) != m + n {return false}

    dp := make([][]bool, m+1)
    for i := range dp {dp[i] = make([]bool, n+1)}

    for i:=0; i<m+1; i++ {
        for j:=0; j< n+1; j++ {
            if i==0 && j==0 {
                dp[i][j] = true
            } else if i == 0 {
                dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1])
            } else if j == 0 {
                dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1])
            } else {
                dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
            }
        }
    }

    return dp[m][n];
}


func main() {
    // INPUT :
    s1 := "aabcc"
    s2 := "dbbca"
    s3 := "aadbbcbcac"

    // OUTPUT :
    fmt.Println(is_interleaving(s1, s2, s3))
}