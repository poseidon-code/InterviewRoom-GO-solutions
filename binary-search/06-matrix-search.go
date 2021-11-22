// PROBLEM: Matrix search
// https://leetcode.com/problems/search-a-2d-matrix/
// https://www.interviewbit.com/problems/matrix-search/

package main

import "fmt"

// SOLUTION
func matrix_search(matrix [][]int, target int) bool {
    if (len(matrix)==0 || len(matrix[0])==0) {return false}

    m, n := len(matrix), len(matrix[0])
    start, end := 0, m*n-1

    for start <= end {
        var mid int = start + (end - start)/2
        var e int = matrix[mid/n][mid%n]

        if target < e {
            end = mid - 1
        } else if target > e {
            start = mid + 1
        } else {
            return true
        }
    }

    return false
}


func main() {
    // INPUT :
    matrix := [][]int{{1,3,5,7}, {10,11,16,20}, {23,30,34,60}}
    target := 3

    // OUTPUT :
    fmt.Println(matrix_search(matrix, target))
}