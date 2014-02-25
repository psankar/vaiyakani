package main

import (
	"fmt"
)

type trie struct {
	key      byte
	count    int
	children map[byte]*trie
	endNode  bool
}

func (t *trie) AddChild(childKeys []byte) (status bool) {

	var child *trie

	i := t
	for _, ch := range childKeys {
		child = i.children[ch]
		if child == nil {
			child = &trie{}
			child.count = 0
			child.children = make(map[byte]*trie)
			child.endNode = false
			i.children[ch] = child
		}
		i.count++
		i = child
	}
	child.count++
	child.endNode = true
	return true
}

func NewTrie() *trie {
	t := &trie{}
	t.key = '/'
	t.endNode = false
	t.count = 0
	t.children = make(map[byte]*trie)

	return t
}

func (t *trie) PrintTrie() {
	fmt.Println("Trie Contents\n===========")
	for k, ch := range t.children {
		printChildren(ch, []byte{k})
	}
}

func printChildren(t *trie, parent []byte) {
	for k, ch := range t.children {
		if ch.endNode {
			fmt.Printf("%s%c\n", string(parent), k)
		}
		printChildren(ch, append(parent, k))
	}
}

func main() {
	t := NewTrie()
	t.AddChild([]byte("God"))
	t.AddChild([]byte("is"))
	t.AddChild([]byte("Great"))
	t.AddChild([]byte("baba"))
	t.AddChild([]byte("baby"))
	t.PrintTrie()
}
