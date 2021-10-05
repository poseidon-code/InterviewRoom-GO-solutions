// PROBLEM: Circular tour
// https://leetcode.com/problems/gas-station/
// https://practice.geeksforgeeks.org/problems/circular-tour/1

package main

import "fmt"

// SOLTUION
func circular_tour(gas []int, cost []int) int {
	start, total, tank := 0, 0, 0

	for i:=0; i<len(gas); i++ {
		tank += gas[i]-cost[i]
		if tank < 0 {
			start = i+1
			total += tank
			tank = 0
		}
	}

	if total+tank < 0 {return -1} else {return start}
}


func main() {
    // INPUT :
    gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}

    // OUTPUT :
    fmt.Println(circular_tour(gas, cost))
}