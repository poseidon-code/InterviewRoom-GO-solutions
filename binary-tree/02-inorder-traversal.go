// PROBLEM: Inorder traversal
// https://www.interviewbit.com/problems/inorder-traversal/

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
func inorder(root *Node) []interface{} {
    var s []*Node;
    var result []interface{}

    if root==nil {return result}

    c := root.left
    s = append(s, root)

    for len(s)!=0 || c!=nil {
        for c!=nil {
            s = append(s, c)
            c=c.left
        }

        temp := s[len(s)-1]
        s = s[:len(s)-1]
        result = append(result, temp.data)
        c = temp.right
    }

    return result
}


func main() {
	// INPUT :
	elements := []E{{1},{6},{2},{nil},{nil},{3},{nil}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    result := inorder(root)
    fmt.Print("[")
    for _, x := range result {
        if x==nil {continue}
        fmt.Print(x," ")
    }
	fmt.Println("\b]")
}