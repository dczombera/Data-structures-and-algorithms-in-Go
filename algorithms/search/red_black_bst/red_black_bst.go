package red_black_bst

import "errors"

const red bool = true
const black bool = false

type RedBlackBST struct {
	Root *Node
}

type Node struct {
	Key         Key
	Val         Value
	Left, Right *Node
	Color       bool
	Size        int
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

func NewEmptyRedBlackBST() RedBlackBST {
	return RedBlackBST{}
}

func NewRedBlackBST(n *Node) RedBlackBST {
	return RedBlackBST{n}
}

func (bst *RedBlackBST) Size() int {
	return bst.size(bst.Root)
}

func (bst *RedBlackBST) size(n *Node) int {
	if n == nil {
		return 0
	}

	return n.Size
}

func (bst *RedBlackBST) IsEmpty() bool {
	return bst.Size() == 0
}

func (bst *RedBlackBST) Contains(k Key) bool {
	_, err := bst.Get(k)
	if err != nil {
		return false
	}

	return true
}

func (bst *RedBlackBST) Get(k Key) (Value, error) {
	return bst.get(bst.Root, k)
}

func (bst *RedBlackBST) get(n *Node, k Key) (Value, error) {
	for n != nil {
		cmp := k.CompareTo(n.Key)
		if cmp < 0 {
			n = n.Left
		} else if cmp > 0 {
			n = n.Right
		} else {
			return n.Val, nil
		}
	}

	return Value(""), errors.New("key not found")
}

func (bst *RedBlackBST) Put(k Key, v Value) {
	bst.Root = bst.put(bst.Root, k, v)
}

func (bst *RedBlackBST) put(n *Node, k Key, v Value) *Node {
	if n == nil {
		return &Node{k, v, nil, nil, red, 1}
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

	if isRed(n.Right) {
		n = bst.rotateLeft(n)
	}
	if isRed(n.Left) && isRed(n.Left.Left) {
		n = bst.rotateRight(n)
	}
	if isRed(n.Left) && isRed(n.Right) {
		flipColors(n)
	}
	return n
}

// ***********************************
// 		Rotation helper functions
// ***********************************
func (bst *RedBlackBST) rotateLeft(n *Node) *Node {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = red
	x.Size = n.Size
	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return x
}

func (bst *RedBlackBST) rotateRight(n *Node) *Node {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = red
	x.Size = n.Size
	n.Size = 1 + bst.size(n.Left) + bst.size(n.Right)
	return x
}

func flipColors(n *Node) {
	if !isRed(n.Left) || !isRed(n.Right) {
		return
	}
	n.Color = red
	n.Left.Color = black
	n.Right.Color = black
}

func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return n.Color == red
}
