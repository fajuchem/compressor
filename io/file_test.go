package io

import (
	"bytes"
	"os"
	"testing"
)

func TestWriteGood(t *testing.T) {
	file := "good"
	defer os.Remove(file)
	input := []byte{48, 48, 49, 49, 48, 48, 48, 49}

	inputWritten := Write(file, input)
	output := Read(file)

	if !bytes.Equal(inputWritten, output) {
		t.Error("Input written different from read")
	}

}
