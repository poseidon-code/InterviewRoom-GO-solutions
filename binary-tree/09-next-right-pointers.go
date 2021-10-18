// PROBLEM: Populating next right pointers
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

package main

type E struct {
	e interface{}
}

type Node struct {
	data interface{}
	left *Node
	right *Node
    next *Node
}

func NBT(elements []E, root *Node, i int, n int) *Node {
	if n==0 { return nil }

	if i<n {
		temp := Node{
			data: elements[i].e,
			left: nil,
			right: nil,
            next: nil,
		}
		root = &temp
		root.left = NBT(elements, root.left, 2*i+1, n)
		root.right = NBT(elements, root.right, 2*i+2, n)
	}

	return root
}


// SOLUTION
func connect(root *Node) {
    for root!=nil && root.left!=nil {
        c := root
        for c!=nil {
            c.left.next = c.right
            if c.next!=nil {
                c.right.next = c.next.left
            } else {
                c.right.next = nil
            }
            c = c.next
        }
        root = root.left
    }
}


func main() {
	// INPUT :
	elements := []E{{1},{2},{3},{4},{5},{6},{7}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    connect(root)
}