// PROBLEM: Vertical order traversal
// https://www.interviewbit.com/problems/vertical-order-traversal-of-binary-tree/

package main

import (
	"fmt"
	"sort"
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
var m [][]int
func traverse(root *Node, x int, y int) {
    if root!=nil {
        m = append(m, []int{x, y, root.data.(int)})
        traverse(root.left, x-1, y-1)
        traverse(root.right, x+1, y-1)
    }
}

func vertical(root *Node) [][]int {
    traverse(root, 0, 0)
    sort.Slice(m, func(i, j int) bool {
        return m[i][0] < m[j][0] || m[i][0] == m[j][0] && m[i][1] > m[j][1] || m[i][0] == m[j][0] && m[i][1] == m[j][1] && m[i][2] < m[j][2]
    })
    var result [][]int

    temp := []int{m[0][2]}
    x := m[0][0]

    for i:=1; i<len(m); i++ {
        if m[i][0] != x {
            t := append([]int{}, temp...)
            result = append(result, t)
            temp = []int{m[i][2]}
            x = m[i][0]
        } else {
            temp = append(temp, m[i][2])
        }
    }

    return append(result, temp)
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3},{4},{5},{6},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(vertical(root))
}