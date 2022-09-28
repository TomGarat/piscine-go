// Write a function ListLast that returns the Data interface of the last element of a linked list l.

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

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return l.Head.Data
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, "3")
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

func ListPushBack(l *List, data string) {
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
