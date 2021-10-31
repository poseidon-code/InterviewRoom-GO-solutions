// PROBLEM: Maximum path sum between two leaf nodes
// https://www.codingninjas.com/codestudio/problems/maximum-path-sum-between-two-leaves_794950

package main

import (
	"fmt"
	"math"
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
func traverse(root *Node, result int) int {
    if root==nil {return 0}
    if root.left==nil && root.right==nil {return root.data.(int)}

    ls := traverse(root.left, result)
    rs := traverse(root.right, result)

    if root.left!=nil && root.right!=nil {
        result = int(math.Max(float64(result), float64(ls+rs+root.data.(int))))
        return int(math.Max(float64(ls), float64(rs))) + root.data.(int)
    }

    return func()int{if root.left==nil {return rs+root.data.(int)}else{return ls+root.data.(int)}}()
}

func leaf_path_sum(root *Node) int {
    result := math.MinInt32
    v := traverse(root, result)

    if result==math.MinInt32 {return v}
    return result
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(leaf_path_sum(root))
}