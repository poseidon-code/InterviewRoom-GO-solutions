// PROBLEM: Level order traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/

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
func levelorder(root *Node) [][]interface{} {
    var q []*Node;
    var result [][]interface{}
    var c []interface{}

    if root==nil {return result}

    q = append(q, root)
    q = append(q, nil)

    for len(q)!=0 {
        t := q[0]
        q = q[1:]

        if t==nil {
            result = append(result, c)
            c = make([]interface{}, 0)

            if len(q) > 0 {q = append(q, nil)}
        } else {
            c = append(c, t.data)
            if t.left!=nil {q = append(q, t.left)}
            if t.right!=nil {q = append(q, t.right)}
        }
    }

    return result
}


func main() {
	// INPUT :
	elements := []E{{3},{9},{20},{nil},{nil},{15},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    result := levelorder(root)
    fmt.Print("[")
    for _, x := range result {
        fmt.Print("[")
        for _, y := range x {
            if y==nil {continue}
            fmt.Print(y,",")
        }
        fmt.Print("\b],")
    }
	fmt.Println("\b]")
}