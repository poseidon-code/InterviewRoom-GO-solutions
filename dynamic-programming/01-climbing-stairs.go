// PROBLEM: Climbing stairs
// https://leetcode.com/problems/climbing-stairs/

package main

import "fmt"

// SOLUTION
func climb_stairs(n int) int {
    a, b := 1, 1
    for i:=n; i>0; i-- {
        a, b = b, a+b
    }
    return a
}


func main() {
    // INPUT :
    n := 3

    // OUTPUT :
    fmt.Println(climb_stairs(n))
}