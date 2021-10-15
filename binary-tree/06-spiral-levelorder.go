// PROBLEM: Level order traversal in sipral form
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// https://www.interviewbit.com/problems/zigzag-level-order-traversal-bt/

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
func levelorder_spiral(root *Node) [][]interface{} {
    var q []*Node;
    var result [][]interface{}
    l2r := true

    if root==nil {return result}

    q = append(q, root)

    for len(q)!=0 {
        s := len(q)
        row := make([]interface{}, s)

        for i:=0; i<s; i++ {
            t := q[0]
            q = q[1:]

            index := func()int{if l2r {return i} else {return s-1-i}}()

            row[index] = t.data
            if t.left!=nil {q = append(q, t.left)}
            if t.right!=nil {q = append(q, t.right)}
        }

        l2r = !l2r
        result = append(result, row)
    }

    return result
}


func main() {
	// INPUT :
	elements := []E{{3},{9},{20},{nil},{nil},{15},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    result := levelorder_spiral(root)
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