// PROBLEM: Delete node from BST
// https://leetcode.com/problems/delete-node-in-a-bst/

package main

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


// SOLUTION
func delete_node(root *Node, key int) *Node {
    if root==nil {return nil}

    if key<root.data.(int) {
        root.left = delete_node(root.left, key)
    } else if key>root.data.(int) {
        root.right = delete_node(root.right, key)
    } else {
        if root.right==nil {
            return root.left
        } else if root.left==nil {
            return root.right
        }

        temp := root.right
        for temp.left!=nil {
            temp = temp.left
        }
        temp.left = root.left
        return root.right
    }

    return root
}


func main() {
    // INPUT :
    elements := []E{{5},{3},{6},{2},{4},{nil},{7}}
    key := 3
    root := NBST(elements, &Node{data: elements[0].e}, 0, len(elements))

    // OUTPUT :
    delete_node(root, key)
}