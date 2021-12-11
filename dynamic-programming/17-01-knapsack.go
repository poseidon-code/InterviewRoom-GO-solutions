// PROBLEM: 01 Knapsack
// https://practice.geeksforgeeks.org/problems/0-1-knapsack-problem/0

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func knapsack(W int, weights []int, values []int, N int) int {
    dp := make([]int, W+1)

    for i:=1; i<N+1; i++ {
        for w:=W; w>=0; w-- {
            if weights[i-1] <= w {
                dp[w] = int(math.Max(float64(dp[w]), float64(dp[w - weights[i-1]] + values[i-1])))
            }
        }
    }

    return dp[W]
}


func main() {
    // INPUT :
    N := 3
    W := 4
    values := []int{1,2,3}
    weights := []int{4,5,1}

    // OUTPUT :
    fmt.Println(knapsack(W, weights, values, N))
}