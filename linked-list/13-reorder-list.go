// PROBLEM: Reorder list
// https://leetcode.com/problems/reorder-list/
// https://www.interviewbit.com/problems/reorder-list/

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
func reorder(head *Node) *Node {
	if head==nil || head.next==nil {
		return head
	}

	var prev, slow, fast, l1, l2 *Node
	prev, slow, fast, l1, l2 = nil, head, head, head, nil

	for fast!=nil && fast.next!=nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}
	prev.next = nil

	var p, c, n *Node
	p, c, n = nil, slow, nil
	for c!=nil {
		n=c.next
		c.next=p
		p=c
		c=n
	}
	l2 = p

	for l1!=nil {
		n1, n2 := l1.next, l2.next
		l1.next = l2

		if n1==nil {
			break
		}

		l2.next=n1
		l1=n1
		l2=n2
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
	result := reorder(head)
	ll.print_ll(result)
}