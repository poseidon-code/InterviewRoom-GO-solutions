// PROBLEM: Check if 2 binary trees are mirror of each other
// https://practice.geeksforgeeks.org/problems/check-mirror-in-n-ary-tree1528/1

package main

import "fmt"

// SOLUTION
func check_mirror(n int, e int, A []int, B []int) bool {
    mp := make(map[int][]int)

    for i:=0; i<2*e; i+=2 {
        mp[A[i]] = append(mp[A[i]], A[i+1])
    }

    for i:=0; i<2*e; i+=2 {
        if B[i+1] != mp[B[i]][len(mp[B[i]])-1] {
            return false
        }
        mp[B[i]] = mp[B[i]][:len(mp[B[i]])-1]
    }

    return true
}


func main() {
	// INPUT :
    n := 3
    e := 2
    A := []int{1,2,1,3}
    B := []int{1,3,1,2}

	// OUTPUT :
    fmt.Println(check_mirror(n, e, A, B))
}