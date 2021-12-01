// PROBLEM: Print longest common subsequence
// https://www.hackerrank.com/challenges/dynamic-programming-classics-the-longest-common-subsequence/problem

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func lcs(A []int, B []int) []int {
    n, m := len(A), len(B)
    lca := make([][]int, n+1)
    for i:=range lca {
        lca[i] = make([]int, m+1)
    }
    
    for i:=0; i<=n; i++ {
        for j:=0; j<=m; j++ {
            if i==0 || j==0 {
                lca[i][j] = 0
            } else if A[i-1]==B[j-1] {
                lca[i][j] = lca[i-1][j-1] + 1
            } else {
                lca[i][j] = int(math.Max(float64(lca[i][j-1]), float64(lca[i-1][j])))
            }
        }
    }

    i, j, k := n, m, lca[n][m]
    var C []int = make([]int, k)
    for i>0 && j>0 {
        if A[i-1]==B[j-1] {
            C[k-1] = A[i-1]
            k--
            i--
            j--
        } else {
            if lca[i-1][j]>lca[i][j-1] {
                i--
            } else {
                j--
            }
        }
    }

    return C
}


func main() {
    // INPUT :
    A := []int{1,2,3,4,1}
    B := []int{3,4,1,2,1,3}

    // OUTPUT :
    fmt.Println(lcs(A, B))
}