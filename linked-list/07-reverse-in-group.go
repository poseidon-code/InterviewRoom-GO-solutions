// PROBLEM: Reverse list in a group of given size k
// https://leetcode.com/problems/reverse-nodes-in-k-group/
// https://www.interviewbit.com/problems/k-reverse-linked-list/

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
func reverse(first *Node, last *Node) *Node {
	p := last
	for first!=last {
		t := first.next
		first.next = p
		p = first
		first = t
	}
	return p
}

func reverse_group(head *Node, k int) *Node {
	node := head

	for i:=0; i<k; i++ {
		if node==nil {
			return head
		}
		node = node.next
	}

	new_head := reverse(head, node)
	head.next = reverse_group(node, k)

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
	reversed := reverse_group(head, 2)
	ll.print_ll(reversed)
}