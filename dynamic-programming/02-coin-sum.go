// PROBLEM: Coin sum infinite
// https://www.interviewbit.com/problems/coin-sum-infinite/

package main

import "fmt"

// SOLUTION
func coin_sum(s []int, n int) int {
    dp := make([]int, n+1)
    dp[0] = 1

    for i:=0; i<len(s); i++ {
        for j:=s[i]; j<=n; j++ {
            dp[j] += dp[j-s[i]]
        }
    }

    return dp[n]
}


func main() {
    // INPUT :
    s := []int{1,2,3}
    n := 4

    // OUTPUT :
    fmt.Println(coin_sum(s, n))
}