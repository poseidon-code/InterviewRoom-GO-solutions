// PROBLEM: Check if a binary tree is BST
// https://practice.geeksforgeeks.org/problems/check-for-bst/1

package main

import (
	"fmt"
)

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
func traverse(root *Node, min *Node, max *Node) bool {
    if root==nil {return true}
    if min!=nil && root.data.(int)<=min.data.(int) {return false}
    if max!=nil && root.data.(int)>=max.data.(int) {return false}

    l := traverse(root.left, min, root)
    r := traverse(root.right, root, max)

    return l && r
}

func is_bst(root *Node) bool {
    return traverse(root, nil, nil)
}


func main() {
	// INPUT :
	elements := []E{{2},{1},{3}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(is_bst(root))
}