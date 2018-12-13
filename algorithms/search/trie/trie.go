package trie

import "errors"

const sizeAlphabet = 256

type Trie struct {
	root *Node
}

type Node struct {
	next  []*Node
	value *Value
}

type Value int

var initCap = 8

func Constructor() Trie {
	return Trie{}
}

func (this Trie) Get(key string) (int, error) {
	n := this.get(this.root, key, 0)
	if n == nil || n.value == nil {
		return -1, errors.New("Key not found")
	}
	return int(*n.value), nil
}

func (this Trie) get(n *Node, key string, pos int) *Node {
	if n == nil {
		return nil
	}
	if len(key) == pos {
		return n
	}
	return this.get(n.next[key[pos]], key, pos+1)
}

func (this *Trie) Put(key string, value int) {
	this.root = this.put(this.root, key, value, 0)
}

func (this *Trie) put(node *Node, key string, value int, pos int) *Node {
	if node == nil {
		node = &Node{make([]*Node, sizeAlphabet), nil}
	}

	if len(key) == pos {
		val := Value(value)
		node.value = &val
	} else {
		node.next[key[pos]] = this.put(node.next[key[pos]], key, value, pos+1)
	}
	return node
}

func (this Trie) Keys() []string {
	return this.KeysWithPrefix("")
}

func (this Trie) KeysWithPrefix(pre string) []string {
	q := make([]string, 0, initCap)
	this.collect(this.get(this.root, pre, 0), pre, &q)
	return q
}

func (this Trie) collect(n *Node, pre string, q *[]string) {
	if n == nil {
		return
	}
	if n.value != nil {
		*q = append(*q, pre)
	}
	for i, c := range n.next {
		this.collect(c, pre+string(i), q)
	}
}

func (this Trie) KeysThatMatch(pattern string) []string {
	q := make([]string, 0, initCap)
	this.collectPattern(this.root, "", pattern, &q)
	return q
}

func (this Trie) collectPattern(n *Node, pre, pattern string, q *[]string) {
	if n == nil {
		return
	}
	if len(pattern) == len(pre) && n.value != nil {
		*q = append(*q, pre)
	}
	if len(pattern) == len(pre) {
		return
	}

	char := pattern[len(pre)]
	for i, c := range n.next {
		if char == '.' || byte(i) == char {
			this.collectPattern(c, pre+string(i), pattern, q)
		}
	}
}

func (this Trie) LongestPrefixOf(s string) string {
	length := this.search(this.root, s, 0, 0)
	return s[0:length]
}

func (this Trie) search(n *Node, s string, pos, length int) int {
	if n == nil {
		return length
	}
	if n.value != nil {
		length = pos
	}
	if pos == len(s) {
		return length
	}
	char := n.next[s[pos]]
	return this.search(char, s, pos+1, length)

}
