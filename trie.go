package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type result struct {
	word  string
	score int
}

type trie struct {
	key      string
	children map[string]*trie
	values   []result
}

func (t *trie) AddWord(key string, value result) (status bool) {

	var child *trie

	if value.word == "அடிமைகள்" {
		fmt.Fprintf(os.Stderr, "Got \n")
	}

	iter := t
	for i := 0; i < len(key); i++ {
		char := string(key[i])

		/* Find if there is a child node for this exact char.
		 * This may be the most frequent case based on initial
		 * assumption. This may change after measurement. */
		child = iter.children[char]
		if child != nil {
			iter = child
			continue
		}

		/* Find if any of the patriciated children
		 * begins with the incoming character */
		for node := range iter.children {
			//fmt.Fprintf(os.Stderr, "Key: [%s]\tNode: [%s]\t\n\a", key, node)
			if node[0] == key[i] {
				var j int

				if node == key[i:] {
					iter.children[node].values = append(iter.children[node].values, value)
					return true
				}

				for j = 1; j < len(node) && ((j + i) < len(key)); {
					if node[j] == key[i+j] {
						j++
						continue
					} else {
						/* Create two children for the
						 * node. (1) The elements from j
						 * until end of the current key.
						 * (2) A new node from (i+j) of
						 * the input key to the function
						 * that was passed */

						node2 := &trie{}
						node2.children = make(map[string]*trie, 0)
						node2.values = append(node2.values, iter.children[node].values...)

						node3 := &trie{}
						node3.children = make(map[string]*trie, 0)
						node3.values = append(node3.values, value)

						node1 := &trie{}
						node1.children = make(map[string]*trie, 2)
						//node1.values = nil
						node1.children[node[j:]] = node2
						node1.children[key[(i+j):]] = node3

						/*
													fmt.Fprintf(os.Stderr,
							`Got node: [%s](%s) and key [%s]
							Created children: [%s](%s) [%s](%s)
							& modified the parent [%s]

							`,
													node, iter.values, key, node[j:], node2.values, key[(i+j):], node3.values, node[:j])
						*/

						iter.children[node].values = nil
						delete(iter.children, node)
						iter.children[node[:j]] = node1

						return true
					}
				}

				if len(key[i:]) > len(node) {
					child = &trie{}
					child.children = make(map[string]*trie, 0)
					child.values = append(child.values, value)
					iter.children[node].children[key[i+len(node):]] = child
					//fmt.Fprintf(os.Stderr, "Created a new child node (%s) only parent-key[%s]\n", child.values, key[len(node):])
					return true
				} else if len(key[i:]) < len(node) {
					node2 := &trie{}
					node2.children = make(map[string]*trie, 0)
					node2.values = append(node2.values, iter.children[node].values...)

					node1 := &trie{}
					node1.children = make(map[string]*trie, 1)
					node1.values = append(node1.values, value)
					node1.children[node[len(key[i:]):]] = node2

					iter.children[node].values = nil
					delete(iter.children, node)
					iter.children[key[i:]] = node1
					return true
				} else {
					fmt.Fprintf(os.Stderr, "Key: [%s]\tNode: [%s]\tToken: [%s]\n\a", key, node, key[i:])
					panic("BUG")
				}
			}
		}

		/* Create a new patriciated node, with all the pending
		 * english alphabets as the key */
		child = &trie{}
		child.children = make(map[string]*trie, 0)
		child.values = append(child.values, value)
		iter.children[key[i:]] = child
		return true
	}
	return false
}

func NewTrie() *trie {
	t := &trie{}
	t.key = "/"
	t.children = make(map[string]*trie, 26)

	return t
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

func (t *trie) GetSuggestions(ch string) ([]string, *trie) {
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

	printAsJSON(t, 1)
	/*

		indentLevel := 4
		for k, ch := range t.children {
			for i := 0; i < indentLevel; i++ {
				fmt.Print(" ")
			}
			fmt.Printf("\"%s\": {\n", string(k))

			if len(ch.values) != 0 {
				for _, value := range ch.values {
					for j := 0; j < indentLevel+4; j++ {
						fmt.Print(" ")
					}
					fmt.Printf("\"value\": \"%s\",\n", value.word)
				}
			}
			printAsJSON(ch, indentLevel+4)

			for i := 0; i < indentLevel; i++ {
				fmt.Print(" ")
			}
			fmt.Println("},")
		}

	*/
}

func printAsJSON(t *trie, indentLevel int) {
	for k, ch := range t.children {
		for i := 0; i < indentLevel; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("\"%s\": {\n", string(k))

		if len(ch.values) != 0 {
			for j := 0; j < indentLevel+1; j++ {
				fmt.Print("\t")
			}
			fmt.Printf("\"value\":[")
			for _, str := range ch.values {
				fmt.Printf("\"%s\",", str.word)
			}
			fmt.Println("],")
		}
		printAsJSON(ch, indentLevel+1)

		for i := 0; i < indentLevel; i++ {
			fmt.Print("\t")
		}
		fmt.Println("},")
	}
}

func main() {
	var t *trie

	fmt.Println("{")

	for ch := 'a'; ch <= 'a'; ch++ {
		fmt.Fprintf(os.Stderr, "%c\n", ch)

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
						if rune(englishWord[0]) == ch {
							t.AddWord(string(englishWord), result{tamilWord, score})
						}
					}
					//fmt.Println()
				}
			}

		}

		t.PrintAsJSON()
	}

	fmt.Println("}")
}
