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
type NodeI struct {
	Data int
	Next *NodeI
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}

}
func ListPushBack(l *List, data interface{}) {
	n := &NodeL{
		Data: data,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}

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

func ListSize(l *List) int {
	ptr := l.Head
	count := 0
	for ptr != nil {
		count++
		ptr = ptr.Next
	}
	return count

}
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	return l.Tail.Data
}
func PrintList1(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}
func ListClear(l *List) {
	l.Head = nil
	l.Tail = nil
}
func ListAt(l *NodeL, pos int) *NodeL {
	ptr := l
	for pos > 0 {
		if ptr == nil {
			return nil
		}
		ptr = ptr.Next
		pos--

	}
	return ptr
}
func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	len := ListSize(l)
	last_adresse := l.Tail.Next

	for len > 0 {
		ptr := l.Head
		for ptr.Next != nil && ptr.Next != last_adresse {
			ptr.Data, ptr.Next.Data = ptr.Next.Data, ptr.Data
			ptr = ptr.Next
		}
		last_adresse = ptr
		len--
	}
}

func ListReverse1(l *List) {
	ptr := l.Head
	pre := l.Tail.Next
	l.Tail = l.Head

	for ptr != nil {
		post := ptr.Next
		ptr.Next = pre
		pre = ptr
		ptr = post
	}
	l.Head = pre
}

func ListForEach(l *List, f func(*NodeL)) {
	ptr := l.Head
	for ptr != nil {
		f(ptr)
		ptr = ptr.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	ptr := l.Head
	for ptr != nil {
		if comp(ptr.Data, ref) {
			return &ptr.Data
		}
		ptr = ptr.Next
	}
	return nil

}
func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	for l.Head.Data == data_ref {

		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
			return
		}

	}

	ptr := l.Head

	for ptr != nil {

		if ptr.Next.Data == data_ref {
			if ptr.Next.Next == nil {
				ptr.Next = nil
				break
			} else {
				ptr.Next = ptr.Next.Next
			}
		} else {
			if ptr.Next == nil {
				break
			}
			ptr = ptr.Next
		}
	}
	l.Tail = ptr

}
func ListMerge(l1 *List, l2 *List) {
	if l2.Head == nil {
		return
	}
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	}
	ptr := l2.Head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	l1.Tail.Next = l2.Head
	l1.Tail = ptr

}
func ListSort(l *NodeI) *NodeI {
	ptr := l
	len := 0
	for ptr != nil {
		len++
		ptr = ptr.Next
	}
	
	for len > 0 {
		ptr = l
		for ptr.Next != nil {

			if ptr.Data > ptr.Next.Data {
				ptr.Data, ptr.Next.Data = ptr.Next.Data, ptr.Data
			}
			ptr = ptr.Next
		}
		len--
	}
	return l

}

func SortListInsert(l *NodeI, data_ref int) *NodeI{
	n := &NodeI{Data: data_ref}
	ptr := l 
	if l == nil{
		return n
	}
	if l.Data >= data_ref{
		n.Next = l
		l = n
		return l
	}
	for ptr.Next != nil{
		if ptr.Data < data_ref && ptr.Next.Data >= data_ref{
			n.Next = ptr.Next
			ptr.Next = n
			return l
		}
		ptr =ptr.Next
	}
	ptr.Next = n
	return l


}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	ptr := n2
	for ptr != nil{
		n1 = SortListInsert(n1,ptr.Data)
		ptr = ptr.Next
	}
	return n1

}
func main() {
	var link *NodeI
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))

}
