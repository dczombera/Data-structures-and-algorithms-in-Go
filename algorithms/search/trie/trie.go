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

func Constructor() Trie {
	return Trie{}
}

func (this Trie) Get(key string) (int, error) {
	return this.get(this.root, key, 0)
}

func (this Trie) get(n *Node, key string, pos int) (int, error) {
	if n == nil {
		return -1, errors.New("Key does not exist")
	}
	if len(key) == pos {
		if n.value == nil {
			return -1, errors.New("Key does not exist")
		}
		return int(*n.value), nil
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
