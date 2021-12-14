// PROBLEM: Longest Bitonic Subsequence
// https://www.interviewbit.com/problems/length-of-longest-subsequence/

package main

import (
	"fmt"
)

// SOLUTION
func lbs(A []int) int {
    n := len(A)

    lis := make([]int, n)
    for i := range lis {lis[i]=1}
    for i:=1; i<n; i++ {
        for j:=0; j<i; j++ {
            if A[i]>A[j] && lis[i]<lis[j]+1 {
                lis[i] = lis[j]+1
            }
        }
    }

    lds := make([]int, n)
    for i := range lds {lds[i]=1}
    for i:=n-2; i>=0; i-- {
        for j:=n-1; j>i; j-- {
            if A[i]>A[j] && lds[i]<lds[j]+1 {
                lds[i] = lds[j]+1
            }
        }
    }

    result := lis[0] + lds[0] - 1
    for i:=1; i<n; i++ {
        if lis[i]+lds[i]-1 > result {
            result = lis[i] + lds[i] -1
        }
    }

    return result
}


func main() {
    // INPUT :
    A := []int{1,2,1}

    // OUTPUT :
    fmt.Println(lbs(A))
}