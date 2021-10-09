// PROBLEM: Maximum height of binary tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// https://www.interviewbit.com/problems/max-depth-of-binary-tree/

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
func maximum_height(root *Node) int {
	if root==nil {
		return 0
	} else {
		return int(math.Max(float64(maximum_height(root.left)), float64(maximum_height(root.right)) + 1))
	}
}


func main() {
	// INPUT :
	elements := []E{{3},{9},{20},{nil},{nil},{15},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
	fmt.Println(maximum_height(root))
}