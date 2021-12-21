// PROBLEM: Region in binary matrix
// https://www.interviewbit.com/problems/region-in-binarymatrix/
// https://practice.geeksforgeeks.org/problems/length-of-largest-region-of-1s-1587115620/1

package main

import (
	"fmt"
	"math"
)

// SOLUTION
var result int

func dfs(grid [][]int, i, j int) {
    if i<0 || j<0 || i>=len(grid) || j>=len(grid[0]) || grid[i][j]==0 {return}

    grid[i][j] = 0
    result++
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i,   j+1)
    dfs(grid, i,   j-1)
    dfs(grid, i+1, j+1)
    dfs(grid, i-1, j+1)
    dfs(grid, i+1, j-1)
    dfs(grid, i-1, j-1)
}

func region(grid [][]int) int {
    n, m, maxarea := len(grid), len(grid[0]), 0

    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            result = 0
            if grid[i][j] == 1 {
                dfs(grid, i, j)
                maxarea = int(math.Max(float64(maxarea), float64(result)))
            }
        }
    }

    return maxarea
}


func main() {
    // INPUT :
    grid := [][]int{{1,1,1,0},{0,0,1,0},{0,0,0,1}}

    // OUTPUT :
    fmt.Println(region(grid))
}