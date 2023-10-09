package main

type Trie struct {
	Children [26]*Trie
	isLeaf   bool
}

func Constructor() Trie {

	return Trie{}
}

func (this *Trie) Insert(word string) {

	node := this
	for _, ch := range word {
		if node.Children[ch-'a'] == nil {
			node.Children[ch-'a'] = &Trie{}
		}
		node = node.Children[ch-'a']
	}
	node.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		if node.Children[ch-'a'] == nil {
			return false
		}
		node = node.Children[ch-'a']
	}

	return node.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		if node.Children[ch-'a'] == nil {
			return false
		}
		node = node.Children[ch-'a']
	}
	return true
}
