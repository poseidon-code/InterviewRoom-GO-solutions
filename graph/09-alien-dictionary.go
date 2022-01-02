// PROBLEM: Alien Dictionary
// https://practice.geeksforgeeks.org/problems/alien-dictionary/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph

package main

import (
	"fmt"
	"math"
	"strings"
)

// SOLUTION
var index int
func toposort(adj [][]int, k int) []string {
    index = k-1
    ans := make([]string, k)
    visited := make([]bool, k)

    for i:=0; i<k; i++ {
        if !visited[i] {
            dfs(adj, ans, k, i, visited)
        }
    }

    return ans
}

func dfs(adj [][]int, ans []string, k, i int, visited []bool) {
    visited[i] = true
    childs := adj[i]

    for j:=0; j<len(childs); j++ {
        if !visited[childs[j]] {
            dfs(adj, ans, k, childs[j], visited)
        }
    }
    
    ans[index] = string(rune(i+'a'))
    index--
}

func order(d []string, n, k int) string {
    adj := make([][]int, k)

    for i:=0; i<len(d)-1; i++ {
        s1 := d[i]
        s2 := d[i+1]

        for j:=0; j<int(math.Min(float64(len(s1)), float64(len(s2)))); j++ {
            if s1[j] != s2[j] {
                adj[int(s1[j])-97] = append(adj[int(s1[j])-97], int(s2[j])-97)
                break
            }
        }

    }
    
    ans := toposort(adj, k)
    return strings.Join(ans, "")
}


func main() {
    // INPUT :
    d := []string{"caa", "aaa", "aab"}
    n := 3
    k := 3

    // OUTPUT :
    fmt.Println(order(d, n, k))
}