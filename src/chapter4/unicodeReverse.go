package main
import (
	"fmt"
)
func main() {
	var s = []byte("this is my string lets reverse it!");
	s = reverse(s)
	fmt.Printf("%s\n",s);

}

func reverse(b []byte) []byte {
	out := b[:0]	
	//going to need to copy it
	for j:=len(b)-1; j > 1 ;j-- {
		out = append (out,b[j])
	}
	return out
}