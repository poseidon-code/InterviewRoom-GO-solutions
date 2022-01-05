// PROBLEM: Word Ladder
// https://leetcode.com/problems/word-ladder/
// https://www.interviewbit.com/problems/word-ladder-i/

package main

import (
	"fmt"
)

// SOLUTION
func ladder_length(begin_word, end_word string, word_list []string) int {
    transformationCount := 1
    queue := []string{}
    queue = append(queue, begin_word)
    words := make(map[string]struct{})
    
    visited := make(map[string]bool)
    visited[begin_word] = true
    for _, v := range word_list {
        words[v] = struct{}{}
    }
    
    if _, ok := words[end_word]; !ok {
        return 0
    }
    
    for len(queue) > 0 {
        queueLength := len(queue)
        
        for i := 0; i < queueLength; i++ {
            word := queue[0]
            queue = queue[1:]
            
            for j := 0; j < len(word); j++ {
                for letter := 'a'; letter <= 'z'; letter++ {
                    modified := word[:j] + string(letter) + word[j + 1:]
                    
                    if modified == end_word {
                        return transformationCount + 1
                    }
                    
                    if _, ok := words[modified]; !ok || modified == word {
                        continue
                    }
                    
                    if _, ok := visited[modified]; !ok {
                        queue = append(queue, modified)
                        visited[modified] = true
                    }
                }
            }
        }

        transformationCount++
    }
    
    return 0
}


func main() {
    // INPUT :
    begin_word := "hit"
    end_word := "cog"
    word_list := []string{"hot","dot","dog","lot","log","cog"}

    // OUTPUT :
    fmt.Println(ladder_length(begin_word, end_word, word_list))
}