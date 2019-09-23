package main

import (
	"fmt"
	"github.com/fajuchem/compressor/huffman"
	"github.com/fajuchem/compressor/io"
)

func main() {
	text := "aabbc"
	filename := "test.bin"

	encodedText := huffman.Encode(text)
	fmt.Println("written:", encodedText)

	io.Write(filename, encodedText)
	data, _ := io.Read(filename)
	fmt.Println("read:", data)

	var result []byte
	fmt.Println("all", data)
	for _, b := range data {
		for i := uint(1); i <= 8; i++ {
			t := b & (1 << (8 - i))
			fmt.Println("aqui", t)
			if t > 0 {
				result = append(result, byte('1'))
			} else {
				result = append(result, byte('0'))
			}
		}
	}

	//decoded := huffman.Decode(data)

	//fmt.Println(string(decoded))

	//if encodedText != "this is an example for huffman encoding" {
	//	fmt.Println(encodedText)
	//}
}
