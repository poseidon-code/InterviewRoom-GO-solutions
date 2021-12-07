// PROBLEM: Coin change
// https://leetcode.com/problems/coin-change/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func change(coins []int, amount int) int {
    need := make([]int, amount+1)
    for i := range need { need[i] = amount+1 }
    need[0] = 0

    for _, c := range coins {
        for a:=c; a<=amount; a++ {
            need[a] = int(math.Min(float64(need[a]), float64(need[a-c]+1)))
        }
    }

    return func()int{if need[len(need)-1] > amount {return -1} else {return need[len(need)-1]}}()
}


func main() {
    // INPUT :
    coins := []int{1,2,5}
    amount := 11

    // OUTPUT :
    fmt.Println(change(coins, amount))
}