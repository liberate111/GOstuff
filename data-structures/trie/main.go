package main

import "fmt"

// AlpSize is the number of possible characters in the trie
const AlpSize = 26

// Node represent each node in the trie
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert will take in a work and add it to the trie
func (t *Trie) Insert(w string) {
	curNode := t.root
	for _, v := range w {
		charIdx := v - 'a'
		if curNode.children[charIdx] == nil {
			curNode.children[charIdx] = &Node{}
		}
		curNode = curNode.children[charIdx]
	}
	curNode.isEnd = true
}

// Search will take in a word and return true is that word is included in the tri
func (t *Trie) Search(w string) bool {
	curNode := t.root
	for _, v := range w {
		charIdx := v - 'a'
		if curNode.children[charIdx] == nil {
			return false
		}
		curNode = curNode.children[charIdx]
	}
	if curNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("wizard"))
	fmt.Println(myTrie.Search("oreo"))
}
