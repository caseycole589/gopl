package main

import (
	"fmt"
)


func main() {
	st := []string{"one","one","one","two","three","four","five","five","six","sevevn","seven"}
	st = removeDuplicates(st)	
	fmt.Printf("%s\n",st)
}

func removeDuplicates(s []string) []string{
	out := s[:0]	
	for i:=0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			out = append(out,s[i])
		}
	}
	if s[len(s)-2] != s[len(s)-1] {
		out = append(out, s[len(s)-1])
	}
	return out
}