package main

import (
	"fmt"
)
type NodeI struct {
	Data int
	Next *NodeI
}
func aSortListInsert(l *NodeI, data_ref int) *NodeI{
	listPushBack(l,data_ref)
	if l == nil{
		return nil
	}
	flag := true
	for flag{
		flag = false
		current := l
		for current.Next != nil{
			if current.Data > current.Next.Data {
				current.Data , current.Next.Data = current.Next.Data , current.Data 
				flag = true
			}
			current =current.Next
		}
	}
	return l
}
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode

	return l
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

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}