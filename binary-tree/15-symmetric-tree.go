// PROBLEM: Symmetric tree
// https://leetcode.com/problems/symmetric-tree/
// https://www.interviewbit.com/problems/symmetric-binary-tree/

package main

import "fmt"

type E struct {
	e interface{}
}

type Node struct {
	data interface{}
	left *Node
	right *Node
}

func NBT(elements []E, root *Node, i int, n int) *Node {
	if n==0 { return nil }

	if i<n {
		temp := Node{
			data: elements[i].e,
			left: nil,
			right: nil,
		}
		root = &temp
		root.left = NBT(elements, root.left, 2*i+1, n)
		root.right = NBT(elements, root.right, 2*i+2, n)
	}

	return root
}


// SOLUTION
func traverse(p *Node, q *Node) bool {
    if p==nil && q==nil {return true
    } else if p==nil || q==nil {return false}

    if p.data != q.data {return false}

    return traverse(p.left, q.right) && traverse(p.right, q.left)
}

func is_symmetric(root *Node) bool {
    if root==nil {return true}
    return traverse(root.left, root.right)
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{2},{3},{4},{4},{3}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(is_symmetric(root))
}