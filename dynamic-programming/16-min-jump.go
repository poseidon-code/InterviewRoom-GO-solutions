// PROBLEM: Minimum jump to reach end
// https://leetcode.com/problems/jump-game-ii/
// https://www.interviewbit.com/problems/min-jumps-array/

package main

import (
	"fmt"
	"math"
)

// SOLUTION
func jump(nums []int) int {
    jumps := 0
    ce, cf := 0, 0

    for i:=0; i<len(nums)-1; i++ {
        cf = int(math.Max(float64(cf), float64(i+nums[i])))

        if i == ce {
            jumps++
            ce = cf
        }
    }

    return jumps
}


func main() {
    // INPUT :
    nums := []int{2,3,1,1,4}

    // OUTPUT :
    fmt.Println(jump(nums))
}