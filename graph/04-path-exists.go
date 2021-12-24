// PROBLEM: Find whether the pat exists
// https://www.interviewbit.com/problems/path-in-directed-graph/
// https://practice.geeksforgeeks.org/problems/find-whether-path-exist5238/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph

package main

import "fmt"

// SOLUTION
func path(grid [][]int, i, j int) bool {
    if i<0 || j<0 || i>=len(grid) || j>=len(grid[i]) || grid[i][j]==0 {return false}
    if grid[i][j]==2 {return true}
    grid[i][j] = 0

    return path(grid, i+1, j) || path(grid, i-1, j) || path(grid, i, j+1) || path(grid, i, j-1)
}

func is_possible(grid [][]int) bool {
    for i:=0; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            if grid[i][j] == 1 {
                return path(grid, i, j)
            }
        }
    }

    return false
}


func main() {
    // INPUT :
    grid := [][]int{{1,3},{3,2}}

    // OUTPUT :
    fmt.Println(is_possible(grid))
}