// PROBLEM: All Root to Leaf path sum
// https://leetcode.com/problems/path-sum-ii/
// https://www.interviewbit.com/problems/root-to-leaf-paths-with-sum/

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
func traverse(root *Node, target int, path []int, paths *[][]int) {
    if root==nil {return}
    path = append(path, root.data.(int))

    if root.left==nil && root.right==nil && target==root.data.(int) {
        *paths = append(*paths, path)
    } else {
        traverse(root.left, target - root.data.(int), path, paths)
        traverse(root.right, target - root.data.(int), path, paths)
    }

    path = path[:len(path)-1]
}

func all_path_sum(root *Node, target int) [][]int {
    paths :=  [][]int{}
    path := []int{}

    traverse(root, target, path, &paths)
    return paths
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{2}}
    target := 3
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    fmt.Println(all_path_sum(root, target))
}