package node

type Node struct {
	Item Item
	Next *Node
}

type Item int
