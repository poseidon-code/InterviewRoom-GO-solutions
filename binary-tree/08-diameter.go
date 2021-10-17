// PROBLEM: Diameter of binary tree
// https://leetcode.com/problems/diameter-of-binary-tree/

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
var d = 0
func traverse(root *Node) int {
    if root==nil {return 0}
    ld := traverse(root.left)
    rd := traverse(root.right)
    d = int(math.Max(float64(d), float64(ld+rd)))
    return int(math.Max(float64(ld), float64(rd)))+1
}

func diameter(root *Node) int {
    traverse(root)
    return d
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3},{4},{5}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(diameter(root))
}