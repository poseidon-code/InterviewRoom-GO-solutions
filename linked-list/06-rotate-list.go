// PROBLEM: Rotate a linked list
// https://leetcode.com/problems/rotate-list/
// https://www.interviewbit.com/problems/rotate-list/

package main

import "fmt"

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
func rotate(head *Node, k int) *Node {
	if head==nil {
		return head
	}

	l := 1
	new_head, tail := head, head

	for tail.next!=nil {
		tail = tail.next
		l++
	}
	tail.next = head

	if k%l == k {
		for i:=0; i<l-k; i++ {
			tail = tail.next
		}
	}

	new_head = tail.next
	tail.next = nil

	return new_head
}


func main() {
	// INPUT :
	var ll LinkedList;
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	ll.add(5)
	head := ll.get_head()

	// OUTPUT :
	rotated := rotate(head, 2)
	ll.print_ll(rotated)
}