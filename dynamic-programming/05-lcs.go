// PROBLEM: Longest common subsequence (delete operation for two strings)
// https://leetcode.com/problems/delete-operation-for-two-strings/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func lcs(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i:=range dp {
        dp[i] = make([]int, n+1)
    }

    for i:=m; i>=0; i-- {
        for j:=n; j>=0; j-- {
            if i<m || j<n {
                dp[i][j] = 
                    func()int{
                        if i<m && j<n && word1[i]==word2[j] {
                            return dp[i+1][j+1]
                        } else {
                            return 1 + int(math.Min(
                                    float64(func()int{if i<m {return dp[i+1][j]} else {return math.MaxInt}}()),
                                    float64(func()int{if j<n {return dp[i][j+1]} else {return math.MaxInt}}()),
                                ))
                        }
                    }()
            }
        }
    }

    return dp[0][0]
}


func main() {
    // INPUT :
    word1 := "sea"
    word2 := "eat"

    // OUTPUT :
    fmt.Println(lcs(word1, word2))
}