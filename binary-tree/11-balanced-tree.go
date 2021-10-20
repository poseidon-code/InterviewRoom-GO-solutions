// PROBLEM: Check if binary tree is balanced
// https://leetcode.com/problems/balanced-binary-tree/

package main

import (
	"fmt"
	"math"
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
func traverse(root *Node) int {
    if root==nil {return 0}

    return int(math.Max(float64(traverse(root.left)), float64(traverse(root.right))))+1
}

func is_balanced(root *Node) bool {
    if root==nil {return true}

    l := traverse(root.left)
    r := traverse(root.right)

    return math.Abs(float64(l-r)) <= 1 && is_balanced(root.left) && is_balanced(root.right)
}


func main() {
	// INPUT :
	elements := []E{{9},{3},{20},{nil},{nil},{15},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(is_balanced(root))
}