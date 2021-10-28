// PROBLEM: Bottom view of binary tree
// https://practice.geeksforgeeks.org/problems/bottom-view-of-binary-tree/1

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
type pair struct {
    n *Node
    l int
}

func bottom_view(root *Node) []int {
    var minlevel, maxlevel int = 0, 0;

    var result []int
    var q []pair
    var m map[int]int = make(map[int]int)
    q = append(q, pair{root, 0})

    for len(q)!=0 {
        x := q[0].n
        p := q[0].l
        q = q[1:]

        if _, ok := m[p]; !ok {
            m[p] = x.data.(int)
            if minlevel>p {minlevel = p}
            if maxlevel<p {maxlevel = p}
        }
        if (x.left!=nil) { q = append(q, pair{x.left, p-1}) }
        if (x.right!=nil) { q = append(q, pair{x.right, p+1}) }
    }

    for i:=minlevel; i<=maxlevel; i++ {
        if _, ok := m[i]; ok {
            result = append(result, m[i])
        }
    }

    return result
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(bottom_view(root))
}