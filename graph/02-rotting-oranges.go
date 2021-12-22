// PROBLEM: Rotting Oranges
// https://leetcode.com/problems/rotting-oranges/
// https://practice.geeksforgeeks.org/problems/rotten-oranges/0

package main

import (
	"fmt"
)

// SOLUTION
func rot(grid [][]int, i, j, d int) int {
    if i<0 || j<0 || i>=len(grid) || j>=len(grid[i]) || grid[i][j]!=1 {return 0}
    grid[i][j] = d + 3
    return 1
}

func oranges(grid [][]int) int {
    fresh, d := 0, 0

    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[i]); j++ {
            if grid[i][j]==1 {fresh++}
        }
    }

    old:=fresh
    for fresh>0 {
        for i:=0; i<len(grid); i++ {
            for j:=0; j<len(grid[i]); j++ {
                if grid[i][j] == d+2 {
                    fresh -= rot(grid, i+1, j, d) +
                             rot(grid, i-1, j, d) +
                             rot(grid, i, j+1, d) +
                             rot(grid, i, j-1, d)
                }
            }
        }
        
        if fresh==old {return -1}
        d++
        old=fresh
    }

    return d
}


func main() {
    // INPUT :
    grid := [][]int{{2,1,1},{1,1,0},{0,1,1}}

    // OUTPUT :
    fmt.Println(oranges(grid))
}