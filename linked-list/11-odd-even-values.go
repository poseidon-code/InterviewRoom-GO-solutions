// PROBLEM: Segregate even and odd valued nodes in a linked list
// https://www.geeksforgeeks.org/segregate-even-and-odd-elements-in-a-linked-list/

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
func oddeven_values(head *Node) *Node {
	if head==nil || head.next==nil {
		return head
	}

	var es, ee, os, oe *Node
	c := head

	for c!=nil {
		val := c.data

		if val%2==0 {
			if es==nil {
				es = c
				ee = es
			} else {
				ee.next = c
				ee = ee.next
			}
		} else {
			if os==nil {
				os = c
				oe = os
			} else {
				oe.next = c
				oe = oe.next
			}
		}

		c = c.next
	}

	if os==nil || es==nil {
		return nil
	}

	ee.next = os
	oe.next = nil

	return es
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
	result := oddeven_values(head)
	ll.print_ll(result)
}