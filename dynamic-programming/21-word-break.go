// PROBLEM: Word Break
// https://leetcode.com/problems/word-break/
// https://www.interviewbit.com/problems/word-break/

package main

import (
	"fmt"
)

// SOLUTION
func word_break(s string, wordDict []string) bool {
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    dict := make(map[string]struct{})
    
    for _, v := range wordDict {
        dict[v] = struct{}{}
    }
    
    for i:=1; i<=n; i++ {
        for j:=0; j<i; j++ {
            if dp[j] {
                substr := s[j:i]
                if _, ok := dict[substr]; ok {
                    dp[i] = true
                    break
                }
            }
        }
    }
    
    return dp[len(dp)-1]
}


func main() {
    // INPUT :
    s := "leetcode"
    wordDict := []string{"leet", "code"}

    // OUTPUT :
    fmt.Println(word_break(s, wordDict))
}