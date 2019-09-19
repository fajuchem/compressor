package huffman

import (
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

type tree interface {
	Freq() int
}

type leaf struct {
	freq  int
	value rune
}

type node struct {
	freq        int
	left, right tree
}

func (n node) Freq() int {
	return n.freq
}

func (l leaf) Freq() int {
	return l.freq
}

type heapTree []tree

func (ht heapTree) Len() int {
	return len(ht)
}

func (ht heapTree) Swap(i, j int) {
	ht[i], ht[j] = ht[j], ht[i]
}

func (ht *heapTree) Pop() (popped interface{}) {
	popped = (*ht)[len(*ht)-1]
	*ht = (*ht)[:len(*ht)-1]
	return
}

func (ht heapTree) Less(i, j int) bool {
	return ht[i].Freq() < ht[j].Freq()
}

func (ht *heapTree) Push(e interface{}) {
	*ht = append(*ht, e.(tree))
}

func buildTree(runesFreq map[rune]int) tree {
	var trees heapTree

	for c, f := range runesFreq {
		trees = append(trees, leaf{f, c})
	}

	heap.Init(&trees)

	for trees.Len() > 1 {
		a := heap.Pop(&trees).(tree)
		b := heap.Pop(&trees).(tree)

		heap.Push(&trees, node{a.Freq() + b.Freq(), a, b})
	}

	return heap.Pop(&trees).(tree)
}

func buildDictionary(t tree, prefix []rune, dic map[rune]string) {
	switch i := t.(type) {
	case leaf:
		dic[i.value] = string(prefix)
	case node:
		prefix = append(prefix, '0')
		buildDictionary(i.left, prefix, dic)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		buildDictionary(i.right, prefix, dic)
		prefix = prefix[:len(prefix)-1]
	}
}

var encodedTree []rune

func encodeTree(t tree) {
	switch i := t.(type) {
	case leaf:
		encodedTree = append(encodedTree, '1')
		encodedTree = append(encodedTree, i.value)
	case node:
		encodedTree = append(encodedTree, '0')
		encodeTree(i.left)
		encodeTree(i.right)
	}
}

func decodeTree(b *bytes.Buffer) tree {
	c, _ := b.ReadByte()

	if c == byte('1') {
		v, _ := b.ReadByte()
		return leaf{1, rune(v)}
	} else {
		left := decodeTree(b)
		right := decodeTree(b)

		return node{1, left, right}
	}

	//if b.UnreadRune() == '1' {
	//	a := []rune(rest[:1])
	//	return leaf{1, a[0]}
	//} else {
	//	left := decodeTree(rest)
	//	right := decodeTree(rest)
	//
	//	return node{1, left, right}
	//}
}

func Decode(text string) string {
	buf := bytes.NewBufferString(text)
	tree := decodeTree(buf)

	buildDictionary(tree, []rune{}, map[rune]string{})
	//printCodes(tree, []byte{})

	return "a"
}

type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
}

func New(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	return bit, nil
}

func Encode(text string) string {
	runesFreq := make(map[rune]int)

	for _, c := range text {
		runesFreq[c]++
	}

	tree := buildTree(runesFreq)

	//var dic = make(map[rune][]rune)
	buildDictionary(tree, []rune{}, dic2)

	encodeTree(tree)
	//printCodes(tree, []byte{})

	var encodedText string
	for _, c := range text {
		encodedText += dic2[c]
	}
	data := []byte(encodedText)
	write(data)

	return string(encodedText)
}

func write(bytes []byte) {
	file, err := os.Create("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []byte
	var currentByte byte = 0
	var i uint = 0
	for _, b := range bytes {
		if b == 49 {
			//fmt.Println(bytes)
			currentByte |= 1 << (7 - i)
		}
		//fmt.Printf("%08b\n", currentByte)

		i++
		// @TODO corrigir: quando nao fecha 8 bits não adiciona o resto
		if i == 8 {
			i = 0
			result = append(result, currentByte)
			currentByte = 0
		}
	}
	fmt.Println(result)

	_, err = file.Write(result)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() string {
	file, _ := os.Open("test.bin")
	var data []byte
	i, _  := file.Read(data)
	fmt.Println(data, i)
	//r := New(bytes.NewBuffer(data))
	fmt.Println("")
	//for i := 0; i < 1; i++ {
	//	for j := 0; j < 8; j++ {
	//		bit, _ := r.ReadBit()
	//		fmt.Printf("%b[%d] = %t\n", data[i], j, bit)
	//	}
	//}

	return "a"
}

func trimFirstRune(s string) (rune, string) {
	v, i := utf8.DecodeRuneInString(s)
	return v, s[i:]
}

var dic2 = make(map[rune]string)

func printCodes(tree tree, prefix []byte) {
	switch i := tree.(type) {
	case leaf:
		// print out symbol, frequency, and code for this
		// leaf (which is just the prefix)
		dic2[i.value] = string(prefix)
		fmt.Printf("%c\t%d\t%s\n", i.value, i.freq, string(prefix))
	case node:
		// traverse left
		prefix = append(prefix, '0')
		printCodes(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		// traverse right
		prefix = append(prefix, '1')
		printCodes(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}

//func write() {
//	//New a bit stream writer with default 5 byte
//	b := bstream.NewBStreamWriter(1)
//	b2 := bstream.NewBStreamWriter(1)
//
//	//Write 0xa0a0 into bstream
//	b.WriteBit(true)
//	b.WriteBit(false)
//	b.WriteBit(false)
//	b2.WriteBit(true)
//	b2.WriteBit(true)
//	b2.WriteBit(false)
//
//	//Read 4 bit out
//	result, _ := b.ReadBits(8)
//	result2, _ := b2.ReadBits(8)
//    
//    var bytes []byte
//
//    bytes = append(bytes, byte(result))
//    bytes = append(bytes, byte(result2))
//    fmt.Println("writing: ", bytes)
//
//    f, _ := os.Create("teste.bin")
//    defer f.Close()
//
//    _, _ = f.Write(bytes)
//}
//
//func read() {
//    f, _ := os.Open("teste.bin")
//    defer f.Close()
//
//    stats, _ := f.Stat()
//
//    var size int64 = stats.Size()
//    _ = size
//    bytes := make([]byte, size)
//
//    bufr := bufio.NewReader(f)
//    _, _ = bufr.Read(bytes)
//
//    fmt.Println("reading: ", bytes)
//}
