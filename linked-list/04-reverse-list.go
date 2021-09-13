// PROBLEM: Reverse a linked list
// https://leetcode.com/problems/reverse-linked-list/
// https://www.interviewbit.com/problems/reverse-linked-list/

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
func reverse(head *Node) *Node {
	var c *Node
	for head!=nil {
		n := head.next
		head.next = c
		c = head
		head = n
	}

	return c
}


func main() {
	// INPUT :
	var ll LinkedList;
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	ll.add(5)
	ll.add(6)
	head := ll.get_head()

	// OUTPUT :
	reversed := reverse(head)
	ll.print_ll(reversed)
}