package main

import (
	"fmt"
	"github.com/fajuchem/compressor/huffman"
)

func main() {
	text := "aabbc"

	//huffman.Read()

	encodedText := huffman.Encode(text)
	huffman.Write(encodedText)
	huffman.Read()

	//s := toBinaryRunes(encodedText)

	//fmt.Println(encodedText)
	//fmt.Println(s)

	//if encodedText != "this is an example for huffman encoding" {
	//	fmt.Println(encodedText)
	//}
}
