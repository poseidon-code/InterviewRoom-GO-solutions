// PROBLEM: Path sum
// https://www.interviewbit.com/problems/path-sum/
// https://leetcode.com/problems/path-sum/

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
func has_path_sum(root *Node, target int) bool {
    if root==nil {return false}

    if root.left == root.right {return target == root.data}

    return has_path_sum(root.left, target - root.data.(int)) || has_path_sum(root.right, target - root.data.(int))
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3}}
    target := 3
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(has_path_sum(root, target))
}