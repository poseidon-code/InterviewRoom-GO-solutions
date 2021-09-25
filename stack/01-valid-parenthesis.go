// PROBLEM: Valid parenthesis
// https://leetcode.com/problems/valid-parentheses/

package main

import "fmt"

// SOLUTION
func is_valid(s string) bool {
	var parenthesis []rune

	for _, c := range s {
		switch c {
			case '{':
				parenthesis = append(parenthesis, '}')
			case '[':
				parenthesis = append(parenthesis, ']')
			case '(':
				parenthesis = append(parenthesis, ')')
			default:
				if len(parenthesis)==0 || c!=parenthesis[len(parenthesis)-1] {
					return false
				} else {
					parenthesis = parenthesis[:len(parenthesis)-1]
				}
		}
	}

	return len(parenthesis)==0
}


func main() {
	// INPUT :
	s := "()[]{}"

	// OUTPUT :
	fmt.Println(is_valid(s))
}