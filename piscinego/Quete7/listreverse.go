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

func ListReverse(l *List) {
	var prev *NodeL
	var next *NodeL
	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

func main() {
	link := &List{}

	ListPushBack(link, "1")
	ListPushBack(link, "2")
	ListPushBack(link, "3")
	ListPushBack(link, "4")

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
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
