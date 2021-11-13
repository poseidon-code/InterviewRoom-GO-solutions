// PROBLEM: Sort a nearly sorted array
// https://www.geeksforgeeks.org/nearly-sorted-algorithm/

package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


// SOLUTION
func k_sorted(nums []int, k int) []int {
    size := func()int{if len(nums)==k {return k} else {return k+1}}()
    
    pq := &Heap{}
    heap.Init(pq)
    for i:=0; i<size; i++ {
        heap.Push(pq, nums[i])
    }

    index := 0
    for i:=k+1; i<len(nums); i++ {
        nums[index] = heap.Pop(pq).(int)
        heap.Push(pq, nums[i])
        index++
    }

    for len(*pq)>0 {
        nums[index] = heap.Pop(pq).(int)
        index++
    }

    return nums
}

func main() {
    // INPUT :
    nums := []int{2,6,3,12,56,8}
    k := 3

    // OUTPUT :
    s := k_sorted(nums, k)
    fmt.Println(s)
}