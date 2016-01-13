package main
import (
	"fmt"
)

func main() {
	s:= [...]int {0,1,2,3,5,6,7,8,9}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

func reverse( s *[9]int) {
	
	for i,j := 0, len(s) -1; i < j; i,j = i+1,j-1 {
		s[i],s[j] = s[j],s[i]
	}
}