package main

import (
	"fmt"
)

func main() {

	var s  = []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println("s before rotation",s)
	s = rotate(s,2)
	fmt.Println("s after rotating 2 positions",s)
}
func rotate(t []int, places int) []int {
	s := t[:places]
	t = t[places:]
	t = append(t, s...)
	return t	
}