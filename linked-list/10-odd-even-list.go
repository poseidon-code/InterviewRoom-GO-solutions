// PROBLEM: Segregate even and odd positioned nodes in linked list
// https://leetcode.com/problems/odd-even-linked-list/
// https://practice.geeksforgeeks.org/problems/rearrange-a-linked-list/1

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
func oddeven(head *Node) *Node {
	if head==nil || head.next==nil {
		return head
	}

	odd, even := head, head.next

	if even.next==nil {
		return head
	}

	c, e := even.next, even

	for c!=nil {
		odd.next = c
		e.next = c.next
		odd = odd.next
		e = e.next

		if e!=nil {
			c = e.next
		} else {
			break
		}
	}

	odd.next = even 
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
	result := oddeven(head)
	ll.print_ll(result)
}