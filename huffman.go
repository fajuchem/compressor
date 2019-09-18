package huffman

import (
	"bytes"
	"container/heap"
	"fmt"
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

var dic = make(map[rune][]rune)

func buildDictionary(t tree, prefix []rune) {
	switch i := t.(type) {
	case leaf:
		dic[i.value] = prefix
	case node:
		prefix = append(prefix, '0')
		buildDictionary(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		buildDictionary(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}

var encodedText []rune

func encodeTree(t tree) {
	switch i := t.(type) {
	case leaf:
		encodedText = append(encodedText, '1')
		encodedText = append(encodedText, i.value)
	case node:
		encodedText = append(encodedText, '0')
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

	buildDictionary(tree, []rune{})
	printCodes(tree, []byte{})

	return "a"
}

func Encode(text string) string {
	runesFreq := make(map[rune]int)

	for _, c := range text {
		runesFreq[c]++
	}

	tree := buildTree(runesFreq)
	buildDictionary(tree, []rune{})

	encodeTree(tree)
	printCodes(tree, []byte{})

	for _, c := range text {
		encodedText = append(encodedText, dic[c]...)
	}

	return string(encodedText)
}

func trimFirstRune(s string) (rune, string) {
	v, i := utf8.DecodeRuneInString(s)
	return v, s[i:]
}
func printCodes(tree tree, prefix []byte) {
	switch i := tree.(type) {
	case leaf:
		// print out symbol, frequency, and code for this
		// leaf (which is just the prefix)
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
