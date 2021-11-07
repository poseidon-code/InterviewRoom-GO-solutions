// PROBLEM: Inorder successor in BST
// https://practice.geeksforgeeks.org/problems/inorder-successor-in-bst/1

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
func inorder_successor(root, x *Node) *Node {
    if root==nil {return nil}
    var s *Node

    for root!=nil {
        if x.data.(int)>=root.data.(int) {
            root = root.right
        } else {
            s = root
            root = root.left
        }
    }
    return s
}


func main() {
    // INPUT :
    elements := []E{{2},{1},{3}}
    root := NBST(elements, &Node{data: elements[0].e}, 0, len(elements))
    x := root.left

    // OUTPUT :
    fmt.Println(inorder_successor(root, x).data)
}