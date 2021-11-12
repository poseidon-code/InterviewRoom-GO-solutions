// PROBLEM: Find kth largest element in a stream
// https://leetcode.com/problems/kth-largest-element-in-a-stream/

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

type KthLargest struct {
    pq      *Heap
    qsize   int
}

// SOLUTION
func k_largest(k int, nums []int) KthLargest {
    h := &Heap{}
    heap.Init(h)
    for i := 0; i < len(nums); i++ {
        heap.Push(h, nums[i])
    }
    for len(*h) > k {heap.Pop(h)}
    return KthLargest{h, k}
}

func (this *KthLargest) add(val int) int {
    if len(*this.pq) < this.qsize {
        heap.Push(this.pq, val)
    } else if val > (*this.pq)[0] {
        heap.Push(this.pq, val)
        heap.Pop(this.pq)
    }
    return (*this.pq)[0]
}

func main() {
    // INPUT :
    k := 3
    nums := []int{4,5,8,2}
    input := []int{3,5,10,9,4}

    // OUTPUT :
    s := k_largest(k, nums)
    for _, i := range input {
        fmt.Println(s.add(i))
    }
}