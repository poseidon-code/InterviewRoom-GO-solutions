// PROBLEM: Min cost path
// https://www.interviewbit.com/problems/min-sum-path-in-matrix/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func path(A [][]int) int {
    m, n := len(A), len(A[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for i:=m-1; i>=0; i-- {
        for j:=n-1; j>=0; j-- {
            if i==m-1 && j==n-1 {
                dp[i][j] = A[i][j]
            } else if i==m-1 {
                dp[i][j] = A[i][j] + dp[i][j+1]
            } else if j==n-1 {
                dp[i][j] = A[i][j] + dp[i+1][j]
            } else {
                dp[i][j] = A[i][j] + int(math.Min(float64(dp[i+1][j]), float64(dp[i][j+1])))
            }
        }
    }

    return dp[0][0]
}


func main() {
    // INPUT :
    A := [][]int{{1,3,2},{4,3,1},{5,6,1}}

    // OUTPUT :
    fmt.Println(path(A))
}