// PROBLEM: Intersection point in Y shaped linked list
// https://leetcode.com/problems/intersection-of-two-linked-lists/
// https://www.interviewbit.com/problems/intersection-of-linked-lists/

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


func createY(headA *Node, headB *Node, skipA int, skipB int) {
	pA, pB := headA, headB

	for i:=skipA; i>1; i-- {
		pA = pA.next
	}
	for i:=skipB; i>1; i-- {
		pB = pB.next
	}
	pB.next = pA.next
}


// SOLUTION
func intersection(headA *Node, headB *Node) *Node {
	pA, pB := headA, headB

	if pA==nil || pB==nil {
		return nil
	}

	for (pA!=nil && pB!=nil && pA!=pB) {
		pA = pA.next
		pB = pB.next

		if pA==pB {
			return pA
		}
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
	}

	return pA
}


func main() {
	// INPUT :
	var llA, llB LinkedList
	llA.add(4); llA.add(1); llA.add(8); llA.add(4); llA.add(5);
	llB.add(5); llB.add(6); llB.add(1); llB.add(8); llB.add(4); llB.add(5);
	headA := llA.get_head()
	headB := llB.get_head()

	// OUTPUT :
	createY(headA, headB, 2, 3)
	intersected := intersection(headA, headB)
	fmt.Println("Intersected at", intersected.data)
}