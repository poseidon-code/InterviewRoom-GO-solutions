// PROBLEM: Rod cutting
// https://www.geeksforgeeks.org/cutting-a-rod-dp-13/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func cutting(price []int, n int) int {
    result := make([]int, n+1)
    result[0] = 0

    for i:=1; i<=n; i++ {
        max_val := math.MinInt16
        for j:=0; j<i; j++ {
            max_val = int(math.Max(float64(max_val), float64(price[j] + result[i-j-1])))
        }
        result[i] = max_val
    }

    return result[n]
}


func main() {
    // INPUT :
    price := []int{1,5,8,9,10,17,17,20}
    n := 8

    // OUTPUT :
    fmt.Println(cutting(price, n))
}