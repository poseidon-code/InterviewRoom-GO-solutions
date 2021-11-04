// PROBLEM: Insert Node into binary search tree
// https://leetcode.com/problems/insert-into-a-binary-search-tree/

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
func insert(root *Node, value int) *Node {
    c := &root
    for *c!=nil {
        if (*c).data.(int) > value {
            c = &(*c).left
        } else {
            c = &(*c).right
        }
    }
    *c = &Node{data: value}
    return root
}


func main() {
    // INPUT :
    elements := []E{{4},{2},{7},{1},{3}}
    value := 5
    root := NBST(elements, &Node{data: elements[0].e}, 0, len(elements))


    // OUTPUT :
    insert(root, value)
}