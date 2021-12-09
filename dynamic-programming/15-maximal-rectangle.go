// PROBLEM: Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/
// https://www.interviewbit.com/problems/max-rectangle-in-binary-matrix/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func rectangle(matrix [][]rune) int {
    if len(matrix)==0 || len(matrix[0])==0 {return 0}
    m, n, maxArea := len(matrix), len(matrix[0]), 0

    left := make([]int, n)
    right := make([]int, n)
    height := make([]int, n)
    for i := range right {right[i] = n-1}

    for i:=0; i<m; i++ {
        rB := n-1
        for j:=n-1; j>=0; j-- {
            if matrix[i][j] == '1' {
                right[j] = int(math.Min(float64(right[j]), float64(rB)))
            } else {
                right[j] = n-1
                rB = j-1
            }
        }

        lB := 0
        for j:=0; j<n; j++ {
            if matrix[i][j] == '1' {
                left[j] = int(math.Max(float64(left[j]),float64(lB)))
                height[j]++
                maxArea = int(math.Max(float64(maxArea), float64(height[j] * (right[j] - left[j] + 1))))
            } else {
                height[j] = 0
                left[j] = 0
                lB = j+1
            }
        }
    }

    return maxArea
}


func main() {
    // INPUT :
    matrix := [][]rune{
        {'1','0','1','0','0'},
        {'1','0','1','1','1'},
        {'1','1','1','1','1'},
        {'1','0','0','1','0'},
    }

    // OUTPUT :
    fmt.Println(rectangle(matrix))
}