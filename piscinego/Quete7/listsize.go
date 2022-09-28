// Write a function ListSize that returns the number of elements in a linked list l.

package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {

	n := l.Head
	size := 0

	for n != nil {
		size++
		n = n.Next
	}
	return size
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

func ListPushFront(l *List, data interface{}) {
	var node NodeL = NodeL{Data: data}

	if l.Head == nil {
		l.Head = &node
	} else {
		node.Next = l.Head
		l.Head = &node
	}
}


unc ListPushBack(l *List, data string) {
	var node NodeL
	node.Data = data
	if l.Head == nil {
		l.Head = &node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &node
	}
}