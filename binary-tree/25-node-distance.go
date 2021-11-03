// PROBLEM: Minimum distance between two nodes
// https://practice.geeksforgeeks.org/problems/min-distance-between-two-given-nodes-of-a-binary-tree/1

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

func traverse(root, a *Node) int {
    if root==nil {return 0}
    if root.data==a.data {return 1}
    c := traverse(root.left, a)
    d := traverse(root.right, a)
    if c==0 && d==0 {return 0} else {return c+d+1}
}


// SOLUTION
func distance(root, p, q *Node) int {
    n := lca(root, p, q)

    c := traverse(n, p)
    d := traverse(n, q)

    return c+d-2
}


func main() {
	// INPUT :
	elements := []E{{3},{5},{1},{6},{2},{0},{8},{nil},{nil},{7},{4}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))
    p := root.left
    q := root.right

	// OUTPUT :
    fmt.Println(distance(root, p, q))
}