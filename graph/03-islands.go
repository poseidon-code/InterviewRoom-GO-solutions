// PROBLEM: Number of islands
// https://leetcode.com/problems/number-of-islands/
// https://practice.geeksforgeeks.org/problems/find-the-number-of-islands/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph

package main

import (
	"fmt"
)

// SOLUTION
func dfs(grid [][]rune, i, j int) int {
    if i<0 || j<0 || i>=len(grid) || j>=len(grid[i]) || grid[i][j]=='0' {return 0}
    
    grid[i][j] = '0'
    dfs(grid, i+1, j)
    dfs(grid, i-1, j)
    dfs(grid, i, j+1)
    dfs(grid, i, j-1)
    return 1
}

func islands(grid [][]rune) int {
    islands := 0

    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            islands += dfs(grid, i, j)
        }
    }

    return islands
}


func main() {
    // INPUT :
    grid := [][]rune{
        {'1','1','0','0','0'},
        {'1','1','0','0','0'},
        {'0','0','1','0','0'},
        {'0','0','0','1','1'},
    }

    // OUTPUT :
    fmt.Println(islands(grid))
}