// PROBLEM: Convert binary tree to its mirror tree
// https://practice.geeksforgeeks.org/problems/mirror-tree/1

package main

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
func mirror(root *Node) {
    if root==nil {
        return
    } else {
        var t *Node
        mirror(root.left)
        mirror(root.right)
        t = root.left
        root.left = root.right
        root.right = t
    }
}


func main() {
	// INPUT :
	elements := []E{{1},{3},{2},{nil},{nil},{5},{4}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    mirror(root)
}