package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
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

func (t *trie) PrintAsJSON() {
	fmt.Println("{")
	indentLevel := 4
	for k, ch := range t.children {
		for i := 0; i < indentLevel; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("\"%s\": {\n", string(k))

		if len(ch.values) != 0 {
			for j := 0; j < indentLevel+4; j++ {
				fmt.Print(" ")
			}
			fmt.Printf("\"value\": \"%s\",\n", ch.values[0].word)
		}
		printAsJSON(ch, indentLevel+4)

		for i := 0; i < indentLevel; i++ {
			fmt.Print(" ")
		}
		fmt.Println("},")
	}
	fmt.Println("}")
}

func printAsJSON(t *trie, indentLevel int) {
	for k, ch := range t.children {
		for i := 0; i < indentLevel; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("\"%s\": {\n", string(k))

		if len(ch.values) != 0 {
			for j := 0; j < indentLevel+4; j++ {
				fmt.Print(" ")
			}
			fmt.Printf("\"value\": \"%s\",\n", ch.values[0].word)
		}
		printAsJSON(ch, indentLevel+4)

		for i := 0; i < indentLevel; i++ {
			fmt.Print(" ")
		}
		fmt.Println("},")
	}
}

var t, trieRoot *trie

func main() {

	t = NewTrie()
	for _, engTransliteratedFile := range os.Args[1:] {
		fileBytes, err := ioutil.ReadFile(engTransliteratedFile)
		if err != nil {
			fmt.Println("Error opening file:")
			fmt.Println(err)
			continue
		}

		lines := strings.Split(string(fileBytes), "\n")

		/* note that this will work with only
		 * linux style file line endings */
		for _, line := range lines {
			if line != "" {
				contents := strings.Split(line, ",")
				score, _ := strconv.Atoi(contents[0])
				tamilWord := contents[1:][0]
				englishWords := contents[2:]

				//fmt.Print(score, tamilWord)
				for _, englishWord := range englishWords {
					//fmt.Print(englishWord, " ")
					t.AddWord([]byte(englishWord), result{tamilWord, score})
				}
				//fmt.Println()
			}
		}

	}

	trieRoot = t
	trieRoot.PrintAsJSON()
}

func testTrie() {
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
