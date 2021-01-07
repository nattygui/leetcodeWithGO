package trie

// Trie 节点
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

// Constructor Initialize your data structure here.
func Constructor() Trie {
	return Trie{
		isEnd: false,
		next:  [26]*Trie{},
	}
}

// Insert Inserts a word into the trie.
func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.next[c-'a'] == nil {
			newNode := Constructor()
			node.next[c-'a'] = &newNode
		}
	}
	node.isEnd = true
}

// Search Returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	node := t
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.next[c-'a'] == nil {
			return false
		}
		node = node.next[c-'a']
	}
	return node.isEnd
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if node.next[c-'a'] == nil {
			return false
		}
		node = node.next[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
