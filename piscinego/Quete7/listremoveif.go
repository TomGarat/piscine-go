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

func ListRemoveIf(l *List, data_ref interface{}) {
	it := l.Head
	for it != nil {
		if it.Data == data_ref {
			if it == l.Head {
				l.Head = it.Next
			} else {
				it = it.Next
			}
		} else {
			it = it.Next
		}
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, "1")
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, "1")
	ListPushBack(link, "Hello")

	ListPushBack(link, "1")
	ListPushBack(link, "There")
	ListPushBack(link, "1")
	ListPushBack(link, "1")
	ListPushBack(link, "How")
	ListPushBack(link, "1")
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, "1")
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
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
