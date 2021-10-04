// PROBLEM: First non-repeating character in a stream
// https://practice.geeksforgeeks.org/problems/first-non-repeating-character-in-a-stream1216/1

package main

import "fmt"

// SOLUTION
func non_repeating(a string) string {
	count := make([]int, 26)
	var q []rune
	var result string

	for _, x := range a {
		q = append(q, x)
		count[x-'a']++

		if len(q)!=0 && count[q[0]-'a']==1 {
			result += string(q[0])
		} else if count[q[0]-'a']>1 {
			for len(q)!=0 && count[q[0]-'a']>1 {
				q = q[1:]
			}

			if len(q)!=0 { 
				result += string(q[0]) 
			} else {
				result += "#"
			}
		}
	}

	return result
}


func main() {
    // INPUT :
    a := "aabc"

    // OUTPUT :
    fmt.Println(non_repeating(a))
}