// PROBLEM: Lowest common ancestor in BST
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

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

func NBST(elements []E, root *Node, i int, n int) *Node {
    if n==0 { return nil }

    if i<n {
        temp := Node{
            data: elements[i].e,
            left: nil,
            right: nil,
        }
        root = &temp
        root.left = NBST(elements, root.left, 2*i+1, n)
        root.right = NBST(elements, root.right, 2*i+2, n)
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
    elements := []E{{6},{2},{8},{0},{4},{7},{9},{nil},{nil},{3},{5},{nil},{nil},{nil},{nil}}
    root := NBST(elements, &Node{data: elements[0].e}, 0, len(elements))
    p := root.left
    q := root.left.right

    // OUTPUT :
    fmt.Println(lca(root, p, q).data)
}