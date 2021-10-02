// PROBLEM: Generate binary numbers from 1 to n
// https://www.geeksforgeeks.org/interesting-method-generate-binary-numbers-1-n/

package main

import "fmt"

// SOLUTION
func generate(n int) {
    var q []string
    q = append(q, "1")

    for ; n > 0; n-- {
        s1 := q[0]
        q = q[1:]
        fmt.Print(s1, " ")

        s2 := s1

        q = append(q, s1+"0")
        q = append(q, s2+"1")
    }
    fmt.Println()
}


func main() {
    // INPUT :
    n := 10

    // OUTPUT :
    generate(n)
}