package io

import (
	"bufio"
	"os"
)

func ByteToBit(bytes []byte) []byte {
	var output []byte
	var currentByte byte = 0
	var i uint = 0
	for _, b := range bytes {
		if b == 49 {
			currentByte |= 1 << (7 - i)
		}

		i++
		// @TODO corrigir: quando nao fecha 8 bits não adiciona o resto
		if i == 8 {
			i = 0
			output = append(output, currentByte)
			currentByte = 0
		}
	}

	return output
}

func Write(name string, bytes []byte) error {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func Read(name string) ([]byte, error) {
	f, err := os.Open(name)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	stats, _ := f.Stat()

	var size int64 = stats.Size()
	_ = size
	bytes := make([]byte, size)

	bufr := bufio.NewReader(f)
	_, err = bufr.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
