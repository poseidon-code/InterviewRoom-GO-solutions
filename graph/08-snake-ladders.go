// PROBLEM: Snakes and Ladders
// https://leetcode.com/problems/snakes-and-ladders/
// https://www.interviewbit.com/problems/snake-ladder-problem/

package main

import "fmt"

// SOLUTION
func calc(row, next int) []int {
    x, y := (next-1)/row, (next-1)%row
    if x%2==1 {y = row-1-y}
    x = row-1-x
    return []int{x,y}
}

func snl(board [][]int) int {
    r := len(board)
    q := []int{1}
    step := 0

    for len(q)!=0 {
        n := len(q)

        for i:=0; i<n; i++ {
            t := q[0]; q = q[1:]
            if t==r*r {return step}

            for j:=1; j<=6; j++ {
                next_step := t+j
                if next_step>r*r {break}

                v := calc(r, next_step)
                row, col := v[0], v[1]

                if board[row][col]!=-1 {
                    next_step = board[row][col]
                }

                if board[row][col]!=r*r+1 {
                    q = append(q, next_step)
                    board[row][col] = r*r+1
                }
            }
        }

        step++
    }

    return -1
}


func main() {
    // INPUT :
    board := [][]int{{-1,-1},{-1,3}}

    // OUTPUT :
    fmt.Println(snl(board))
}