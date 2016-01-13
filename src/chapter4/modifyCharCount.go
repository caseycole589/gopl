package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)
func main() {
	counts := make(map[rune]int) 	//counts of unicode chars
	letter := make(map[rune]int)
	digit  := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int 	// counts of length of utf8 encoding
	invalid := 0					//counts of invalid

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()//returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr,"charcount: %v\n",err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letter[r]++
		}
		if unicode.IsDigit(r) {
			digit[r]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\nletters\tcount\n")
	for c, n := range letter {
		fmt.Printf("%q\t%d\n",c, n)
	}
	fmt.Printf("\ndigits\tcount\n")
	for c, n := range digit{
		fmt.Printf("%q\t%d\n",c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}