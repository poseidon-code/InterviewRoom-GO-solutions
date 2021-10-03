// PROBLEM: Minimum time to rot all oranges
// https://leetcode.com/problems/rotting-oranges/
// https://practice.geeksforgeeks.org/problems/rotten-oranges/0

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func rotting(grid [][]int) int {
	var q [][]int
	dir := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
	ct, result := 0, -1

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] > 0 { ct++ }
			if grid[i][j] == 2 { q = append(q, []int{i, j}) }
		}
	}

	for len(q)!=0 {
		result++
		size := len(q)

		for k:=0; k<size; k++ {
			cur := q[0]
			ct--
			q = q[1:]

			for i:=0; i<4; i++ {
				x := cur[0]+dir[i][0]
				y := cur[1]+dir[i][1]

				if x>=len(grid) || x<0 || y>=len(grid[0]) || y<0 || grid[x][y]!=1 {
					continue
				}

				grid[x][y] = 2
				q = append(q, []int{x, y})
			}
		}
	}

	if ct==0 { return int(math.Max(0, float64(result))) }
	return -1
}


func main() {
    // INPUT :
    grid := [][]int{{2,1,1},{1,1,0},{0,1,1}}

    // OUTPUT :
    fmt.Println(rotting(grid))
}