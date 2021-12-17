// PROBLEM: Palindrome Partitioning
// https://leetcode.com/problems/palindrome-partitioning-ii/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func min_cut(s string) int {
    n := len(s)
    cut := make([]int, n+1)
    for i := range cut {cut[i] = i-1}

    for i:=0; i<n; i++ {
        for j:=0; i-j>=0 && i+j<n && s[i-j]==s[i+j]; j++ {
            cut[i+j+1] = int(math.Min(float64(cut[i+j+1]), float64(1 + cut[i-j])))
        }

        for j:=1; i-j+1>=0 && i+j<n && s[i-j+1]==s[i+j]; j++ {
            cut[i+j+1] = int(math.Min(float64(cut[i+j+1]), float64(1 + cut[i-j+1])))
        }
    }

    return cut[n];
}


func main() {
    // INPUT :
    s := "aab"

    // OUTPUT :
    fmt.Println(min_cut(s))
}