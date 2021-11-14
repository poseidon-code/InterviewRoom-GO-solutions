// PROBLEM: Find kth smallest element in matrix
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

package main

import (
	"container/heap"
	"fmt"
)

type point struct {
    matrix *[][]int
    n int
    m int
}

type Heap []point

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return (*(h[i].matrix))[h[i].n][h[i].m] < (*(h[j].matrix))[h[j].n][h[j].m] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(point)) }
func (h *Heap) Pop() interface{} {
    l := len(*h)
    res := (*h)[l - 1]
    *h = (*h)[:l - 1]
    return res
}


// SOLUTION
func k_smallest(matrix [][]int, k int) int {
    n := len(matrix)
    if n == 0 {
        return 0
    }
    
    h := &Heap{}
    heap.Push(h, point{&matrix, 0, 0})
    for i := 1; i < n; i++ {
        heap.Push(h, point{&matrix, i, 0})
        heap.Push(h, point{&matrix, 0, i})
    }
    
    var res point
    
    set := make(map[point]struct{})
    
    for k > 0 {
        res = heap.Pop(h).(point)
        k--
        
        var newpoint point
        if res.n == res.m {
            continue
        } else if res.n > res.m {
            newpoint = point{&matrix, res.n, res.m + 1}
        } else {
            newpoint = point{&matrix, res.n + 1, res.m}
        }
    
        if _, ok := set[newpoint]; !ok {
            set[newpoint] = struct{}{}
            heap.Push(h, newpoint)
        }
    }

    return matrix[res.n][res.m]
}

func main() {
    // INPUT :
    matrix := [][]int{{1,5,9}, {10,11,13}, {12,13,15}}
    k := 8

    // OUTPUT :
    fmt.Println(k_smallest(matrix, k))
}