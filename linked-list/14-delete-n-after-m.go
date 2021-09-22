// PROBLEM: Delete N nodes after M nodes of a linked list
// https://practice.geeksforgeeks.org/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/1

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
func delete_nodes(head *Node, m int, n int) *Node {
	c := head
	var t *Node = nil

	for c!=nil {
		for i:=1; i<m && c!=nil; i++ {
			c = c.next
		}

		if c==nil {
			return head
		}

		t = c.next
		for count:=1; count<=n && t!=nil; count++ {
			t = t.next
		}
		c.next = t
		c = t
	}

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
	ll.add(7)
	ll.add(8)
	ll.add(9)
	head := ll.get_head()

	// OUTPUT :
	result := delete_nodes(head, 2, 5)
	ll.print_ll(result)
}