// PROBLEM: Remove nth node from the end
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// https://www.interviewbit.com/problems/remove-nth-node-from-list-end/

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
func remove_nth(head *Node, n int) *Node {
	slow := head
	fast := head

	for i:=0; i<=n; i++ {
		fast = fast.next
	}

	for fast != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next

	return head
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
	new_head := remove_nth(head, 2)
	ll.print_ll(new_head)
}