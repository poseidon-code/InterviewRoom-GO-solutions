// PROBLEM: Min cost climbing stairs
// https://leetcode.com/problems/min-cost-climbing-stairs/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func min_cost(cost []int) int {
    for i:=2; i<len(cost); i++ {
        cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])))
    }

    return int(math.Min(float64(cost[len(cost)-1]), float64(cost[len(cost)-2])))
}


func main() {
    // INPUT :
    cost := []int{10,15,20}

    // OUTPUT :
    fmt.Println(min_cost(cost))
}