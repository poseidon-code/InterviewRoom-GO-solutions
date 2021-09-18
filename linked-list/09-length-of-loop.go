// PROBLEM: Find length of loop in linked list
// https://practice.geeksforgeeks.org/problems/find-length-of-loop/1

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

func loop(head *Node, p int) {
	if head.next==nil {
		return
	}
	if p==-1 {
		return
	}
	
	tail, node := head, head
	for tail.next!=nil {
		tail = tail.next
	}
	for i:=0; i<p; i++ {
		node = node.next
	}
	
	tail.next = node
}


// SOLUTION
func length_loop(head *Node) int {
	if head==nil || head.next==nil {
		return 0
	}

	slow, fast, entry := head, head, head

	for fast.next!=nil && fast.next.next!=nil {
		slow = slow.next
		fast = fast.next.next
		if slow==fast {
			result := 1
			for slow != entry {
				slow = slow.next
				entry = entry.next
				result++
			}
			return result
		}
	}

	return 0
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
	loop(head, 4)

	// OUTPUT :
	length := length_loop(head)
	fmt.Println(length)
}