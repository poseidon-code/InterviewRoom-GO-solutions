// PROBLEM: Cycle in undirected graph
// https://www.interviewbit.com/problems/cycle-in-undirected-graph/
// https://practice.geeksforgeeks.org/problems/detect-cycle-in-an-undirected-graph/1/?category{]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph

package main

import "fmt"

// SOLUTION
func dfs(adj [][]int, v []bool, p, s int) bool {
    v[s] = true

    for _, i := range adj[s] {
        if v[i]==false {
            if dfs(adj, v, s, i) {
                return true
            }
        } else if i!=p {
            return true
        }
    }

    return false
}

func is_cycle(graph [][]int, n int) bool {
    v := make([]bool, n+1)
    adj := make([][]int, n+1)
    adj[0] = append(adj[0], -1)

    for i:=0; i<len(graph); i++ {
        adj[graph[i][0]] = append(adj[graph[i][0]], graph[i][1])
        adj[graph[i][1]] = append(adj[graph[i][1]], graph[i][0])
    }

    for i:=1; i<n+1; i++ {
        if v[i]==false {
            if dfs(adj, v, -1, i) {
                return true
            }
        }
    }

    return false
}


func main() {
    // INPUT :
    n := 5
    graph := [][]int{{1,2},{1,3},{2,3},{1,4},{4,5}}

    // OUTPUT :
    fmt.Println(is_cycle(graph, n))
}