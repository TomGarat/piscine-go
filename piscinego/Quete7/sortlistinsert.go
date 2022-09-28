package main

import "fmt"

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l == nil {
		return n
	}
	if l.Data > data_ref {
		n.Next = l
		return n
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < data_ref {
		iterator = iterator.Next
	}
	n.Next = iterator.Next
	iterator.Next = n
	return l
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var sorted *NodeI
	for l != nil {
		sorted = SortListInsert(sorted, l.Data)
		l = l.Next
	}
	return sorted
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(ListSort(link))

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(ListSort(link))
}
