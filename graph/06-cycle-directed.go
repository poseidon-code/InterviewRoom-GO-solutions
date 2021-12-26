// PROBLEM: Cycle in a directed graph
// https://leetcode.com/problems/course-schedule/
// https://www.interviewbit.com/problems/cycle-in-directed-graph/

package main

import "fmt"

// SOLUTION
func can_finish(courses int, prerequisites [][]int) bool {
    G := make([][]int, courses)
    degree := make([]int, courses)
    var bfs []int

    for _, e := range prerequisites {
        G[e[1]] = append(G[e[1]], e[0])
        degree[e[0]]++
    }

    for i:=0; i<courses; i++ {
        if degree[i]==0 {bfs = append(bfs, i)}
    }

    for i:=0; i<len(bfs); i++ {
        for _, j := range G[bfs[i]] {
            degree[j]--
            if degree[j] == 0 {
                bfs = append(bfs, j)
            }
        }
    }

    return len(bfs) == courses
}


func main() {
    // INPUT :
    courses := 2
    prerequisites := [][]int{{1,0},{0,1}}

    // OUTPUT :
    fmt.Println(can_finish(courses, prerequisites))
}