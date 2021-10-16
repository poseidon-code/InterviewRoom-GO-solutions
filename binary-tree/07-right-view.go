// PROBLEM: Right view of binary tree
// https://leetcode.com/problems/binary-tree-right-side-view/

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
func right_view(root *Node) []interface{} {
    var q []*Node;
    var result []interface{}

    if root==nil {return result}

    q = append(q, root)

    for len(q)!=0 {
        size := len(q)

        for i:=0; i<size; i++ {
            t := q[0]
            q = q[1:]
            if i==size-1 {result = append(result, t.data)}
            if t.left!=nil {q = append(q, t.left)}
            if t.right!=nil {q = append(q, t.right)}
        }
    }

    return result
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3},{nil},{5},{nil},{4}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    result := right_view(root)
    fmt.Print("[")
    for _, x := range result {
        if x==nil {continue}
        fmt.Print(x," ")
    }
	fmt.Println("\b]")
}