package Structs

type Trie struct {
	treeNode [26]*Trie
	isEnd    bool
}

// MakeTrie /** Initialize your data structure here. */
func MakeTrie() Trie {
	return Trie{}
}

// Insert /** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'A'
		if node.treeNode[ch] == nil {
			node.treeNode[ch] = new(Trie)
		}
		node = node.treeNode[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'A'
		if node.treeNode[ch] == nil {
			return nil
		}
		node = node.treeNode[ch]
	}
	return node
}

// Search /** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
