// PROBLEM: Spiral matrix
// https://leetcode.com/problems/spiral-matrix/
// https://www.interviewbit.com/problems/spiral-order-matrix-i/

package main

import "fmt"

// SOLUTION
func spiral_matrix(matrix [][]int) []int {
	var result []int
	m := len(matrix)
	n := len(matrix[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	steps := []int{n, m - 1}

	id := 0
	ir := 0
	ic := -1

	for steps[id%2] != 0 {
		for i := 0; i < steps[id%2]; i++ {
			ir += directions[id][0]
			ic += directions[id][1]
			result = append(result, matrix[ir][ic])
		}

		steps[id%2]--
		id = (id + 1) % 4
	}

	return result
}

func main() {
	// INPUT :
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	// OUTPUT :
	fmt.Println(spiral_matrix(matrix))
}
