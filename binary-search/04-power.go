// PROBLEM: Power(x,n)
// https://leetcode.com/problems/powx-n/
// https://www.interviewbit.com/problems/implement-power-function/

package main

import "fmt"

// SOLUTION
func power(x float64, n int) float64 {
    if n==0 {return 1}
    if n==1 {return x}
    if n<0 {x, n = 1/x, -n}
    if n%2==0 {return power(x*x, n/2)}

    return power(x*x, n/2) * x
}


func main() {
    // INPUT :
    x := 2.10000
    n := 3

    // OUTPUT :
    fmt.Println(power(x, n))
}