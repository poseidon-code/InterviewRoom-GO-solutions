// PROBLEM: Invert a binary tree
// https://leetcode.com/problems/invert-binary-tree/
// https://www.interviewbit.com/problems/invert-the-binary-tree/

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
func invert(root *Node) *Node {
    var s []*Node
    s = append(s, root)

    for len(s)!=0 {
        p := s[len(s)-1]
        s = s[:len(s)-1]

        if p!=nil {
            s = append(s, p.left)
            s = append(s, p.right)
            p.left, p.right = p.right, p.left
        }
    }

    return root
}


func main() {
	// INPUT :
	elements := []E{{4},{2},{7},{1},{3},{6},{9}}
	root := NBT(elements, &Node{data: elements[0].e}, 0, len(elements))

	// OUTPUT :
    invert(root)
}