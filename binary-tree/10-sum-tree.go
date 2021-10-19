// PROBLEM: Check if sum tree
// https://practice.geeksforgeeks.org/problems/sum-tree/1

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
var result = true
func traverse(root *Node) int {
    if root==nil {return 0}
    if root.left==nil && root.right==nil {return root.data.(int)}

    ld := traverse(root.left)
    rd := traverse(root.right)

    if ld+rd!=root.data {result = false}

    return root.data.(int)
}

func is_sum_tree(root *Node) bool {
    traverse(root)
    return result
}


func main() {
	// INPUT :
	elements := []E{{10},{20},{30},{10},{10}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(is_sum_tree(root))
}