// PROBLEM: Topological Sort (Course Schedule 2)
// https://leetcode.com/problems/course-schedule-ii/

package main

import "fmt"

// SOLUTION
func find_order(courses int, prerequisites [][]int) []int {
    G := make([][]int, courses)
    indegree := make([]int, courses)
    var ans []int

    for _, pre := range prerequisites {
        G[pre[1]] = append(G[pre[1]], pre[0])
        indegree[pre[0]]++
    }

    var q []int
    for i:=0; i<courses; i++ {
        if indegree[i] == 0 {q = append(q, i)}
    }

    for len(q)!=0 {
        cur := q[0]
        q = q[1:]
        ans = append(ans, cur)

        for _, nextCourse := range G[cur] {
            indegree[nextCourse]--
            if indegree[nextCourse] == 0 {
                q = append(q, nextCourse)
            }
        }
    }

    if len(ans) == courses {return ans}
    return []int{}
}


func main() {
    // INPUT :
    courses := 2
    prerequisites := [][]int{{1,0}}

    // OUTPUT :
    fmt.Println(find_order(courses, prerequisites))
}