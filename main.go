package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/node"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/queue"
	"github.com/dczombera/data-structures-and-algorithms-in-go/datastructs/stack"
)

var nToPush = 6
var nToPop = 3

func main() {
	printMsg("Stack")
	s := stack.NewEmptyStack()
	for i := 0; i < nToPush; i++ {
		s.Push(node.Item(i))
	}
	for i := 0; i < nToPop; i++ {
		s.Pop()
	}

	fmt.Printf("\nThere are %d nodes on the stack:\n", s.Size)
	for !s.IsEmpty() {
		item, _ := s.Pop()
		fmt.Printf("%v\n", item)
	}

	printMsg("Queue")
	q := queue.NewEmptyQueue()
	for i := 0; i < nToPush; i++ {
		q.Push(node.Item(i))
	}
	for i := 0; i < nToPop; i++ {
		q.Pop()
	}

	fmt.Printf("\nThere are %d nodes on the queue:\n", q.Size)
	for !q.IsEmpty() {
		item, _ := q.Pop()
		fmt.Printf("%v\n", item)
	}
}

func printMsg(name string) {
	del := fmt.Sprintf("*****%s*******", strings.Repeat("*", utf8.RuneCountInString(name)))
	fmt.Printf("\n%s\n", del)
	fmt.Printf("*\t%s\t*\n", name)
	fmt.Println(del)
}
