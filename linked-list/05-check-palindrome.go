// PROBLEM: Check if a linked list is palindrome
// https://leetcode.com/problems/palindrome-linked-list/
// https://www.interviewbit.com/problems/palindrome-list/

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
func is_palindrome(head *Node) bool {
	var rev *Node
	slow, fast := head, head

	for fast!=nil && fast.next!=nil {
		t:=rev
		fast = fast.next.next
		rev = slow
		slow = slow.next
		rev.next = t
	}

	if fast!=nil {
		slow = slow.next
	}

	for rev!=nil && rev.data==slow.data {
		slow = slow.next
		rev = rev.next
	}

	return rev==nil
}


func main() {
	// INPUT :
	var ll LinkedList;
	ll.add(1)
	ll.add(2)
	ll.add(3)
	ll.add(3)
	ll.add(2)
	ll.add(1)
	head := ll.get_head()

	// OUTPUT :
	fmt.Println(is_palindrome(head))
}