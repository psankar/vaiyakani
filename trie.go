package main

import (
	"fmt"
)

type result struct {
	word  string
	score int
}

type trie struct {
	key      byte
	children map[byte]*trie
	values   []result
}

func (t *trie) AddWord(key []byte, value result) (status bool) {

	var child *trie

	i := t
	for _, ch := range key {
		child = i.children[ch]
		if child == nil {
			child = &trie{}
			child.children = make(map[byte]*trie)
			i.children[ch] = child
		}
		i = child
	}
	child.values = append(child.values, value)
	return true
}

func NewTrie() *trie {
	t := &trie{}
	t.key = '/'
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
		if len(ch.values) != 0 {
			fmt.Printf("%s%c\t", string(parent), k)
			fmt.Println(ch.values)
		}
		printChildren(ch, append(parent, k))
	}
}

func getChildren(t *trie, suggestions *[]string) {
	for _, ch := range t.children {
		for _, value := range ch.values {
			*suggestions = append(*suggestions, value.word)
		}
		getChildren(ch, suggestions)
	}

	if len(t.children) == 0 {
		for _, value := range t.values {
			*suggestions = append(*suggestions, value.word)
		}
	}

}

func (t *trie) GetSuggestions(ch byte) ([]string, *trie) {
	child := t.children[ch]
	if child != nil {
		var suggestions []string
		getChildren(child, &suggestions)
		return suggestions, child
	} else if len(t.children) == 0 {
		var suggestions []string
		for _, value := range t.values {
			suggestions = append(suggestions, value.word)
		}
		return suggestions, t
	} else {
		return nil, t
	}
}

func main() {

	t := NewTrie()
	t.AddWord([]byte("velivantha"), result{"வெளிவந்த", 143})
	t.AddWord([]byte("irukkathu"), result{"இருக்காது", 18})
	t.AddWord([]byte("ithanai"), result{"இத்தனை", 37})
	t.AddWord([]byte("illamale"), result{"இல்லாமலே", 22})
	t.AddWord([]byte("ippadithan"), result{"இப்படித்தான்", 13})
	t.AddWord([]byte("irundhanar"), result{"இருந்தனர்", 12})
	t.PrintTrie()

	trieRoot := t

	inputs := []string{"irukkathuv", "ithanai", "illamale", "velivantha"}
	for _, input := range inputs {

		fmt.Println("====================================================")

		t = trieRoot
		var suggestions []string
		for _, k := range input {
			fmt.Printf("\n%c\n-----------\n", k)
			suggestions, t = t.GetSuggestions(byte(k))
			for _, suggestion := range suggestions {
				fmt.Println(suggestion)
			}
		}
	}
}
