// PROBLEM: kth smallest node in BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

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

var value, result int

func traverse(root *Node) {
    if root==nil || root.data==nil {return}
    traverse(root.left)
    value--
    if value==0 {result = root.data.(int)}
    traverse(root.right)
}

// SOLUTION
func kth_smallest(root *Node, k int) int {
    value = k
    traverse(root)
    return result
}


func main() {
    // INPUT :
    elements := []E{{3},{1},{4},{nil},{2}}
    k := 1
    root := NBST(elements, &Node{data: elements[0].e}, 0, len(elements))

    // OUTPUT :
    fmt.Println(kth_smallest(root, k))
}