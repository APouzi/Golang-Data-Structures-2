package main

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	char     byte
	adajList map[byte]*Node //Remember, nodes are already refrences, we don't need to set pointers to the
	endWord  bool
}

func (trie *Trie) insert(word string) {
	if trie.root == nil {
		trie.root = &Node{char: '#'}
	}

	curr := trie.root
	for i := 0; i < len(word); i++ {
		nextChar := word[i]
		if _, ok := curr.adajList[nextChar]; ok {
			curr = curr.adajList[nextChar]
		} else {
			if curr.adajList == nil {
				curr.adajList = make(map[byte]*Node)
			}
			curr.adajList[nextChar] = &Node{char: nextChar}
			curr = curr.adajList[nextChar]
		}
	}

}

func (trie *Trie) search(word string) {

	curr := trie.root
	for i := 0; i < len(word); i++ {

		nextChar := word[i]
		if _, ok := curr.adajList[nextChar]; ok {
			curr = curr.adajList[nextChar]
			fmt.Println(string(curr.char)) //Put this here because we moved AFTER the root, which is worthless to us.
		} else {
			fmt.Println("Not here!")
			break
		}
	}

}

func main() {
	reTrieval := Trie{}
	reTrieval.insert("apple")
	reTrieval.insert("apples")
	reTrieval.search("apple")
	reTrieval.search("apples")
}
