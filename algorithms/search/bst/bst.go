package bst

import (
	"errors"
	"log"
)

type BST struct {
	Root *Node
}

type Node struct {
	Key   Key
	Val   Value
	Left  *Node
	Right *Node
	Size  int
}

type Key int
type Value string

func (this Key) CompareTo(that Key) int {
	if this > that {
		return 1
	} else if this < that {
		return -1
	}
	return 0
}

func NewEmptyBST() BST {
	return BST{}
}

func NewBST(n *Node) BST {
	return BST{n}
}

func (bst *BST) Size() int {
	return bst.size(bst.Root)
}

func (bst *BST) size(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Size
}

func (bst *BST) Contains(k Key) bool {
	_, error := bst.Get(k)
	if error != nil {
		return false
	}
	return true
}

func (bst *BST) IsEmpty() bool {
	return bst.Size() == 0
}

func (bst *BST) Get(k Key) (Value, error) {
	return bst.get(bst.Root, k)
}

func (bst *BST) get(n *Node, k Key) (Value, error) {
	if n == nil {
		return Value(""), errors.New("Key not found")
	}

	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		return bst.get(n.Left, k)
	} else if cmp > 0 {
		return bst.get(n.Right, k)
	}
	return n.Val, nil
}

func (bst *BST) Put(k Key, v Value) {
	if v == "" {
		bst.Delete(k)
		return
	}
	bst.Root = bst.put(bst.Root, k, v)
}

func (bst *BST) put(n *Node, k Key, v Value) *Node {
	if n == nil {
		return &Node{k, v, nil, nil, 1}
	}

	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		n.Left = bst.put(n.Left, k, v)
	} else if cmp > 0 {
		n.Right = bst.put(n.Right, k, v)
	} else {
		n.Val = v
	}

	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return n

}

func (bst *BST) Delete(k Key) {
	if bst.IsEmpty() {
		log.Fatalln("Tree is empty")
	}
	bst.Root = bst.delete(bst.Root, k)
}

func (bst *BST) delete(n *Node, k Key) *Node {
	if n == nil {
		return nil
	}
	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		n.Left = bst.delete(n.Left, k)
	} else if cmp > 0 {
		n.Right = bst.delete(n.Right, k)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		temp := n
		n := bst.min(temp.Right)
		n.Right = bst.delMin(temp.Right)
		n.Left = temp.Left
	}
	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return n
}

func (bst *BST) DelMin() {
	if bst.IsEmpty() {
		log.Fatalln("Tree is empty")
	}
	bst.Root = bst.delMin(bst.Root)
}

func (bst *BST) delMin(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = bst.delMin(n.Left)
	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return n
}

func (bst *BST) DelMax() {
	if bst.IsEmpty() {
		log.Fatalln("Tree is empty")
	}
	bst.Root = bst.delMax(bst.Root)
}

func (bst *BST) delMax(n *Node) *Node {
	if n.Right == nil {
		return n.Left
	}
	n.Right = bst.delMax(n.Right)
	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return n
}

func (bst *BST) Min() *Node {
	if bst.IsEmpty() {
		log.Fatalln("Tree is empty")
	}
	return bst.min(bst.Root)
}

func (bst *BST) min(n *Node) *Node {
	if n.Left == nil {
		return n
	}
	return bst.min(n.Left)
}

func (bst *BST) Max() *Node {
	if bst.IsEmpty() {
		log.Fatalln("Tree is empty")
	}
	return bst.max(bst.Root)
}

func (bst *BST) max(n *Node) *Node {
	if n.Right == nil {
		return n
	}
	return bst.max(n.Right)
}

func (bst *BST) IsBST() bool {
	return bst.isBST(bst.Root, nil, nil)
}

func (bst *BST) isBST(n *Node, min, max *Key) bool {
	if n == nil {
		return true
	}

	if min != nil && n.Key.CompareTo(*min) <= 0 {
		return false
	}

	if max != nil && n.Key.CompareTo(*max) >= 0 {
		return false
	}

	return bst.isBST(n.Left, min, &n.Key) && bst.isBST(n.Right, &n.Key, max)
}
