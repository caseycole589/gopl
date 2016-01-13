package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	word := make(map[string]int)	//count of string word occurences
	scanner:= bufio.NewScanner(os.Stdin)//input
	scanner.Split(bufio.ScanWords)	//split words
	for scanner.Scan() {
		word[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintf(os.Stderr, "wordcount: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("word\t\tcount\n")
	for i, v := range word {
		fmt.Printf("\n%s\t\t%d\n",i,v)
		fmt.Printf("------------------------------------------")
	}
	
}