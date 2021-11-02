// PROBLEM: Least Common Ancestor of two nodes
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

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
func lca(root, p, q *Node) *Node {
    if root==nil || root==p || root==q {return root}

    left := lca(root.left, p, q)
    right := lca(root.right, p, q)

    return func()*Node{
        if left==nil {
            return right
        } else {
            if right==nil {
                return left
            } else {
                return root
            }
        }
    }()
}


func main() {
	// INPUT :
	elements := []E{{3},{5},{1},{6},{2},{0},{8},{nil},{nil},{7},{4}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))
    p := root.left
    q := root.right

	// OUTPUT :
    fmt.Println(lca(root, p, q).data)
}