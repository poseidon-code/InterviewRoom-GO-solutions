// PROBLEM: Find middle element
// https://leetcode.com/problems/middle-of-the-linked-list/
// https://practice.geeksforgeeks.org/problems/finding-middle-element-in-a-linked-list/1

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
func middle_node(head *Node) *Node {
	slow := head
	fast := head

	for (fast.next != nil) {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}


func main() {
	// INPUT :
	var ll LinkedList;
	ll.add(0)
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(4)
	ll.add(5)
	ll.add(6)
	ll.add(7)
	ll.add(8)
	head := ll.get_head()

	// OUTPUT :
	start := middle_node(head)
	ll.print_ll(start)
}