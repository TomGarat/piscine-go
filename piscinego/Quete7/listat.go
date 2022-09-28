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

func ListAt(l *NodeL, nbr int) *NodeL {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}
	if nbr <= index {
		for l != nil {
			if count == nbr {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, "1")

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
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
