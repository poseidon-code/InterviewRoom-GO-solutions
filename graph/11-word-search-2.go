// PROBLEM: Word Search 2
// https://leetcode.com/problems/word-search-ii/
// https://practice.geeksforgeeks.org/problems/word-boggle4143/1/?category[]=Graph&category[]=Graph&page=2&query=category[]Graphpage2category[]Graph

package main

import (
	"fmt"
)

// SOLUTION
type trienode struct{
    val rune
    word string
    next map[rune]*trienode
}

func find_words(board [][]byte, words []string) []string {
    root := &trienode{next : map[rune]*trienode{}}
    for _, w := range words {
        p := root
        for i, b := range w {
            if _, ok := p.next[b]; !ok {
                p.next[b] = &trienode{val:b, next:map[rune]*trienode{}} 
            }
            if i==len(w)-1 {p.next[b].word = w}
            p = p.next[b]
        }
    }

    result := []string{}
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[i]); j++ {
            dfs(board, &result, root, i,j, map[string]bool{})
        }
    }
    
    return result
}

func dfs(board [][]byte, res *[]string, node *trienode, i,j int, visited map[string]bool){
    k := fmt.Sprintf("%v_%v", i,j)

    if i<0 || j<0 || i>=len(board) || j>=len(board[0]) || visited[k] { return }
    nn, ok:= node.next[rune(board[i][j])]
    if !ok {return}
    
    if len(nn.word)>0 {
        *res = append(*res, nn.word)
        nn.word = ""
    }
    
    visited[k] = true
    dfs(board, res, nn, i+1, j, visited)
    dfs(board, res, nn, i-1, j, visited)
    dfs(board, res, nn, i, j+1, visited)
    dfs(board, res, nn, i, j-1, visited)
    visited[k] = false
}


func main() {
    // INPUT :
    board := [][]byte{
        {'o','a','a','n'},
        {'e','t','a','e'},
        {'i','h','k','r'},
        {'i','f','l','v'},
    }
    words := []string{"oath","pea","eat","rain"}

    // OUTPUT :
    fmt.Println(find_words(board, words))
}