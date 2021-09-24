// PROBLEM: Add two numbers
// https://leetcode.com/problems/add-two-numbers/
// https://www.interviewbit.com/problems/add-two-numbers-as-lists/

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
func add(l1 *Node, l2 *Node) *Node {
	head := &Node{}
	t := head
	c := 0

	for c!=0 || l1!=nil || l2!=nil {
		if l1!=nil {
			c += l1.data
			l1 = l1.next
		}
		if l2!=nil {
			c += l2.data
			l2 = l2.next
		}

		t.next = &Node{data: c%10}
		t = t.next
		c /= 10
	}

	return head.next
}


func main() {
	// INPUT :
	var ll1, ll2 LinkedList;
	ll1.add(2)
	ll1.add(4)
	ll1.add(3)
	ll2.add(5)
	ll2.add(6)
	ll2.add(4)
	h1 := ll1.get_head()
	h2 := ll2.get_head()

	// OUTPUT :
	result := add(h1, h2)
	ll1.print_ll(result)
}