package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	//Write a function ListPushFront that inserts a new element NodeL at the beginning of the list l while using the structure List

	var node NodeL = NodeL{Data: data} // create a new node

	if l.Head == nil { // if the list is empty
		l.Head = &node // the head is the new node
	} else { // if the list is not empty
		node.Next = l.Head // the new node points to the head
		l.Head = &node     // the head is the new node
	}
}

func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
