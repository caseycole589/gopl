package main

import (
	"os"
	"fmt"
	"bytes"
	
)

func main() {
	
	//if no args are given
	if len(os.Args) <= 1	{
		fmt.Println(` Add Commas iteratively:
		Usage: arg1 arg2 ....`)
	} else {
		//iterate through args
		for i := 1; i < len(os.Args); i++ {
			addCommas(os.Args[i])
		}
	}
}

func addCommas(word string) {
	var buf bytes.Buffer
	byteArray := []byte(word)
	found, count := false, 0
		for j := len(byteArray) - 1; j >= 0; j-- {
			if found  {
				count++
				buf.WriteByte(byteArray[j])
				if count % 3 == 0 {
					buf.WriteByte(',')
				}
			} else if(byteArray[j] == '.'){
				found = true
				buf.WriteByte(byteArray[j])	
			} else {
				buf.WriteByte(byteArray[j])
			}
		}

	 reverseBuf(&buf)
	 fmt.Println(&buf)
	 // return buf.String()
}

func reverseBuf(buffer *bytes.Buffer) {
	var newBuf bytes.Buffer
	newByteArr := buffer.Bytes()
	for i := len(newByteArr) -1 ; i >= 0; i-- {
		newBuf.WriteByte(newByteArr[i])
	}
	*buffer = newBuf
}