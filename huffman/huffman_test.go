package huffman_test

import (
	"github.com/fajuchem/compressor/huffman"
	"testing"
)

func TestDecode(t *testing.T) {
	text := "000001x1r1m1n0001c1l1o01h1s00001g1p01t01u1d01e1i001f1a1  1001101111011011111110110111111111100111110110000111110011100010100110111111111011100001111011110011111111111001111110011111011001010010111100111101100110001"

	encodedText := huffman.Decode(text)

	if encodedText != "this is an example for huffman encoding" {
		t.Errorf(encodedText)
	}
}
func TestEncode(t *testing.T) {
	text := "this is an example for huffman encoding"
	//text := "aaaaaabccccccddeeeee"

	encodedText := huffman.Encode(text)

	if encodedText != "111111111111111111111101101101110111" {
		t.Errorf(encodedText)
	}
}
