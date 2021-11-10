// PROBLEM: Find top k frequent elements
// https://leetcode.com/problems/top-k-frequent-elements/

package main

import "fmt"

// SOLUTION
func k_frequent(nums []int, k int) []int {
    counts := make(map[int]int, 0)
    for _, i := range nums {counts[i]++}

    buckets := make([][]int, len(nums)+1)
    for p, q := range counts {buckets[q] = append(buckets[q], p)}

    var result []int
    for i:=len(buckets)-1; i>=0 && len(result)<k; i-- {
        for _, n := range buckets[i] {
            result = append(result, n)
            if len(result)==k {break}
        }
    }

    return result
}

func main() {
    // INPUT :
    nums := []int{1,1,1,2,2,3}
    k := 2

    // OUTPUT :
    fmt.Println(k_frequent(nums, k))
}