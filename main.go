package main

import (
	"fmt"
	"github.com/fajuchem/compressor/huffman"
	"github.com/fajuchem/compressor/io"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := os.Args[1]
	temp, _ := os.Open(inputFile)
	t, _ := ioutil.ReadAll(temp)
	text := string(t[:len(t)-1])
	filename := inputFile + ".bin"

	encodedTree, encodedText := huffman.Encode(text)
	encoded := append(encodedTree, io.ByteToBit(encodedText)...)

	fmt.Println("input:", text)

	io.Write(filename, encoded)
	fmt.Println("written:", encoded)

	data, _ := io.Read(filename)
	fmt.Println("read:", data)

	decoded := huffman.Decode(data)
	fmt.Println("decoded:", decoded)

	fmt.Println("corrupted?", decoded != text)

	//if encodedText != "this is an example for huffman encoding" {
	//	fmt.Println(encodedText)
	//}
}
