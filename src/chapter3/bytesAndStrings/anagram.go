package main

import (
	"os"
	"fmt"
)

func main() {
	
	//if no args are given
	if len(os.Args) != 3 {
		fmt.Println(` Check Anagram 
		Usage: arg1 arg2 `)
	} else {
		fmt.Println(check_annagrams(os.Args[1], os.Args[2]))
	}
}

func check_annagrams(word1 string, word2 string) bool {
	for i, j := 0, len(word2) - 1; i < len(word2) ; i,j = i+1,j-1 {
		if word1[i] != word2[j] {
			return false
		}
	}
	return true
}
