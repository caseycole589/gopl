package main

import (
	"fmt"
	"chapter2/popcount"
)



func main() {

	var s uint64  = 1846744073709551614;
	fmt.Println("s binary representation")
	fmt.Printf("%b\n",s)
	//popcount gives the number of bits that are set 
	//for an unsigned 64 bit int
	//2 4 8 16 32 64
	//number of bits set	
	fmt.Println(popcount.PopCount(s))
}
