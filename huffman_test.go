package huffman_test

import (
	"github.com/fajuchem/compressor"
	"testing"
)

func TestEncode(t *testing.T) {
	text := "aaaabbbbcccddef"

	encodedText := huffman.Encode(text)

	if encodedText != "00010011001111011" {
		t.Errorf(encodedText)
	}
}
