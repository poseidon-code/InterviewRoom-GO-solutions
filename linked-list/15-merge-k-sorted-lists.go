// PROBLEM: Merge k sorted lists
// https://leetcode.com/problems/merge-k-sorted-lists/
// https://www.interviewbit.com/problems/merge-two-sorted-lists/
// https://practice.geeksforgeeks.org/problems/merge-k-sorted-linked-lists/1

package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) add(n int) {
	t := Node{}
	t.data = n
	t.next = nil

	if ll.head == nil {
		ll.head = &t
		ll.tail = &t
	} else {
		ll.tail.next = &t
		ll.tail = ll.tail.next
	}
}

func (ll *LinkedList) print_ll(start *Node) {
	p := start
	fmt.Print("[")
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println("\b]")
}

func (ll *LinkedList) get_head() *Node {
	return ll.head
}


// SOLUTION
type minHeap []*Node
func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].data < h[j].data }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*Node)) }
func (h *minHeap) Pop() interface{} {
    min := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return min
}

func merge_sorted(lists []*Node) *Node {
	pq := &minHeap{}
	heap.Init(pq)
	var head, tail, next *Node

	for i:=0; i<len(lists); i++ {
		if lists[i]!=nil {
			heap.Push(pq, lists[i])
		}
	}

	for pq.Len() > 0 {
		if tail != nil {
			next = heap.Pop(pq).(*Node)
			if next.next != nil {
				heap.Push(pq, next.next)
			}
			tail.next = next
			tail = next
		} else {
			head = heap.Pop(pq).(*Node)
			tail = head
			if head.next != nil {
				heap.Push(pq, head.next)
			}
		}
	}

	return head
}


func main() {
	// INPUT :
	var ll1, ll2, ll3 LinkedList;
	ll1.add(1)
	ll1.add(4)
	ll1.add(5)
	ll2.add(1)
	ll2.add(3)
	ll2.add(4)
	ll3.add(2)
	ll3.add(6)
	h1 := ll1.get_head()
	h2 := ll2.get_head()
	h3 := ll3.get_head()

	lists := []*Node{h1, h2, h3}

	// OUTPUT :
	result := merge_sorted(lists)
	ll1.print_ll(result)
}