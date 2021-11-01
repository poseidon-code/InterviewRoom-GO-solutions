// PROBLEM: Maximum path sum from any node to node
// https://leetcode.com/problems/binary-tree-maximum-path-sum/

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
var max_sum = math.MinInt32

func traverse(root *Node) int {
    if root==nil {return 0}
    c := func()int{if root.data==nil {return 0} else {return root.data.(int)}}()

    ls := traverse(root.left)
    rs := traverse(root.right)
    max_sum = int(math.Max(float64(max_sum), float64(ls+rs+c)))
    return int(math.Max(float64(ls), float64(rs))) + c
}

func node_path_sum(root *Node) int {
    v := traverse(root)

    if max_sum==math.MinInt32 {return v}
    return max_sum
}


func main() {
	// INPUT :
	elements := []E{{-10},{9},{20},{nil},{nil},{15},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(node_path_sum(root))
}