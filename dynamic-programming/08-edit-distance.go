// PROBLEM: Edit distance
// https://leetcode.com/problems/edit-distance/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func edit_distance(word1 string, word2 string) int {
    m, n, pre := len(word1), len(word2), 0
    result := make([]int, n+1)
    for i:=1; i<=n; i++ {result[i] = i}

    for i:=1; i<=m; i++ {
        pre = result[0]
        result[0] = i

        for j:=1; j<=n; j++ {
            temp := result[j]
            if word1[i-1]==word2[j-1] {
                result[j] = pre
            } else {
                result[j] = int(math.Min(float64(pre), math.Min(float64(result[j-1]), float64(result[j])))) + 1
            }
            pre = temp
        }
    }

    return result[n]
}


func main() {
    // INPUT :
    word1 := "horse"
    word2 := "ros"

    // OUTPUT :
    fmt.Println(edit_distance(word1, word2))
}