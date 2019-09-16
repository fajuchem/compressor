package huffman

import (
	"container/heap"
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

var encodedText []rune

func generateString(t tree, prefix []rune) {
	switch i := t.(type) {
	case leaf:
		encodedText = append(encodedText, prefix...)
	case node:
		prefix = append(prefix, '0')
		generateString(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		generateString(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}

func Encode(text string) string {
	runesFreq := make(map[rune]int)

	for _, c := range text {
		runesFreq[c]++
	}

	tree := buildTree(runesFreq)
	generateString(tree, []rune{})

	return string(encodedText)
}
