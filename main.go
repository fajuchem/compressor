package main

import (
	"bytes"
	"fmt"
	"github.com/fajuchem/compressor/huffman"
)

func toBinaryRunes(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		fmt.Fprintf(&buffer, "%b", runeValue)
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}

func main() {
	//text := "aabbc"

	fmt.Println(huffman.Read())

	//encodedText := huffman.Encode(text)

	//s := toBinaryRunes(encodedText)

	//fmt.Println(encodedText)
	//fmt.Println(s)

	//if encodedText != "this is an example for huffman encoding" {
	//	fmt.Println(encodedText)
	//}
}
