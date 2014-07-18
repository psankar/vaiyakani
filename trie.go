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

	iter := t
	for i := 0; i < len(key); i++ {
		char := string(key[i])

		/* Find if there is a child node for this exact char */
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
				/* The first character has matched. So while exiting this block,
				 * We would have created the new node for holding the incoming value.
				 * In case, the key is already existing, we will append the incoming
				 * value to the values that are already part of the key.
				 * single key = multiple values is possible */

				var j int

				/* If the first letter is matching, check if the incoming key is matching the node,
				 * and if so just add the incoming value, as another value for the node */
				if node == key[i:] {
					iter.children[node].values = append(iter.children[node].values, value)
					return true
				}

				for j = 1; j < len(node) && ((j + i) < len(key)); {
					if node[j] == key[i+j] {
						/* Proceed until the next character in the incoming key
						 * and the node are different. At some point they will have to be
						 * different, as otherwise the previous block would have caught it */
						j++
						continue
					} else {
						/* A character has differed between the incoming key and the character in
						 * the node. So we need to split our trie at this point and create new children
						 * to hold the two diverging paths as two different nodes.
						 *
						 * We will create three different nodes:
						 *
						 * node1 - a new node that will have characters from "node" until the
						 * point where it diffes from the key. We will update this to the parent of node
						 *
						 * node2 - a new node that will be the child of node1. It will contain
						 * the characters which were originally in "node" after the position of
						 * difference with the incoming key. The values of the previous "node"
						 * will become the values of this node.
						 *
						 * node3 - a new that will be the child of node1 and will contain the new
						 * incoming value. The suffix that is different from the odl "node" will
						 * be used as the key for identifying this child on node1's children list
						 */

						node2 := &trie{}
						//node2.children = make(map[string]*trie, 0)
						node2.children = iter.children[node].children
						node2.values = append(node2.values, iter.children[node].values...)

						node3 := &trie{}
						node3.children = make(map[string]*trie, 0)
						node3.values = append(node3.values, value)

						node1 := &trie{}
						node1.children = make(map[string]*trie, 2)
						node1.children[node[j:]] = node2
						node1.children[key[(i+j):]] = node3

						iter.children[node].values = nil
						delete(iter.children, node)
						iter.children[node[:j]] = node1

						return true
					}
				}

				if len(key[i:]) > len(node) {
					/* The incoming key is longer than the node string. But all the characters
					 * in the node have matched the characters of the incoming key. So, we will
					 * just create a new child for the node, with the pending suffix in key, as
					 * the identifier for this new child */
					child = &trie{}
					child.children = make(map[string]*trie, 0)
					child.values = append(child.values, value)
					iter.children[node].children[key[i+len(node):]] = child
					//fmt.Fprintf(os.Stderr, "Created a new child node (%s) only parent-key[%s]\n", child.values, key[len(node):])
					return true
				} else if len(key[i:]) < len(node) {
					/* The incoming key is shorter than the node string. But all the characters
					 * in the key have matched the node string. So we need to split the node, at
					 * the last character of the incoming key, and the pending suffix in the node,
					 * i.e, characters beyond len(key) in node will have to become a new node,
					 * and this new node will have all the values that earlier belonged to the
					 * old "node" */
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
					/* A non-reachable condition */
					panic("Unreachable Code")
				}
			}
		}

		/* Create a new node, with all the characters in the incoming key.
		 * This will be executed if none of the first characters in the trie
		 * match the first letter of the incoming key or if the trie is empty */
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

func (t *trie) PrintAsJSON() {
	printAsJSON(t, 1)
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

func (t *trie) SearchForString(input string) string {

	//	fmt.Fprintf(os.Stderr, "\n\nSearching for input string: [%s]\n", input)

	active := t
	result := ""

	i := 0

findLetter:
	for i < len(input) {

		letter := input[i]
		//	fmt.Fprintf(os.Stderr, "Searching for single character [%c]\n", letter)

		if active.children[string(letter)] != nil {
			active = active.children[string(letter)]
			i++
			//	fmt.Fprintf(os.Stderr, "Exact character match for [%c]\n", letter)
			continue
		} else {
			for k, nodes := range active.children {
				//		fmt.Fprintf(os.Stderr, "Checking prefix for string [%s]\n", k)
				if strings.HasPrefix(input[i:], k) {
					//fmt.Fprintf(os.Stderr, "Prefix match found [%s]\n", nodes)
					active = nodes
					i += len(k)
					goto findLetter
				}
			}
			//	fmt.Fprintf(os.Stderr, "Could not find the input string\n")
			return result
		}
	}

	//	fmt.Fprintf(os.Stderr, "Found the input string\n")

	return active.values[0].word
}

func main() {
	var t *trie

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
				//fmt.Fprintf(os.Stderr, "Reading Line [%s]\n", line)
				contents := strings.Split(line, ",")
				score, _ := strconv.Atoi(contents[0])
				tamilWord := contents[1:][0]
				englishWords := contents[2:]

				//fmt.Print(score, tamilWord)
				for _, englishWord := range englishWords {
					//fmt.Print(englishWord, " ")
					t.AddWord(string(englishWord), result{tamilWord, score})
				}
			}
		}
	}

	/* Print Trie
	fmt.Println("{")
	t.PrintAsJSON()
	fmt.Println("}")
	*/
}
