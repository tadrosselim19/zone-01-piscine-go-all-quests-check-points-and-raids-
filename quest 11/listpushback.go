package main

import (
	"fmt"
)
type NodeI struct {
	Data int
	Next *NodeI
}
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	
	if l.Head == nil{
		l.Head = n
		l.Tail =n
	}else{
		l.Tail.Next = n
		l.Tail =n
	}
}
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil{
		l.Head = n
		l.Tail = n
	}else{
		n.Next =l.Head
		l.Head = n
	}

}
func ListLast(l *List) interface{} {
	if l.Head == nil{
		return nil
	}
	var last interface{}
	object:=l.Head 
	for object.Next != nil{	
		object=object.Next
	}
	last = object.Data
	return last
}
func ListClear(l *List) {
	l.Head = nil
	l.Tail = nil
}
func ListSize(l *List) int {
	count := 0
	object := l.Head
	for object != nil {
		count ++
		object =object.Next
	}
	return count
}
func ListAt(l *NodeL, pos int) *NodeL{
	if pos < 0 {
		return nil
	}

	value := l
	count := pos
	for count != 0{
		if value == nil{
			return nil
		}
		value = value.Next
		count--
	}
	return value
}
func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}
func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil{
		f(current)
		current = current.Next
	}
}
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}
func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil{
		if comp(current.Data,ref) {
			return &current.Data
		}
		current=current.Next
	}
	return nil
}
func ListRemoveIf(l *List, data_ref interface{}) {
	for  l.Head != nil && l.Head.Data == data_ref{
		l.Head = l.Head.Next
	}
	if l.Head == nil{
		return
	}
	current := l.Head
	for current.Next != nil{
		if current.Next.Data == data_ref{
			current.Next = current.Next.Next
		}else{
			current = current.Next
		}
		
	}
		l.Tail = current
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}

}

func aListMerge(l1 *List, l2 *List) {
	c1 := l1.Head
	c2 := l2.Head

	if c1 == nil {
		l1.Head = c2
		l1.Tail = l2.Tail
		return
	}

	for c1.Next != nil {
		c1 = c1.Next
	}

	c1.Next = c2
	if l2.Tail != nil {
		l1.Tail = l2.Tail
	}
}
func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	if l2.Head == nil {
		return
	}

	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
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
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
func PrintElem(node *NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *NodeL) {
	node.Data = 2
}
func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Print("nil", "\n")
}
func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
	link2 := &List{}

	ListPushBack(link2, "hello")
	ListPushBack(link2, "how are")
	ListPushBack(link2, "you")
	ListPushBack(link2, 1)

	fmt.Println(ListAt(link2.Head, 3).Data)
	fmt.Println(ListAt(link2.Head, 1).Data)
	fmt.Println(ListAt(link2.Head, 7))

	link3 := &List{}

	ListPushBack(link3, 1)
	ListPushBack(link3, 2)
	ListPushBack(link3, 3)
	ListPushBack(link3, 4)

	ListReverse(link3)

	ita := link.Head

	for ita != nil {
		fmt.Println(ita.Data)
		ita = ita.Next
	}

	fmt.Println("Tail", link3.Tail)
	fmt.Println("Head", link3.Head)

	link4 := &List{}

	ListPushBack(link4, "1")
	ListPushBack(link4, "2")
	ListPushBack(link4, "3")
	ListPushBack(link4, "5")

	ListForEach(link4, Add2_node)

	it := link4.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	link5 := &List{}

	ListPushBack(link5, 1)
	ListPushBack(link5, "hello")
	ListPushBack(link5, 3)
	ListPushBack(link5, "there")
	ListPushBack(link5, 23)
	ListPushBack(link5, "!")
	ListPushBack(link5, 54)

	PrintList(link)

	fmt.Println("--------function applied--------")
	ListForEachIf(link5, PrintElem, IsPositiveNode)

	ListForEachIf(link5, StringToInt, IsAlNode)

	fmt.Println("--------function applied--------")
	PrintList(link5)

	fmt.Println()

	link6 := &List{}

	ListPushBack(link6, "hello")
	ListPushBack(link6, "hello1")
	ListPushBack(link6, "hello2")
	ListPushBack(link6, "hello3")

	found := ListFind(link6, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)



	link7 := &List{}
	link8:= &List{}

	fmt.Println("----normal state----")
	ListPushBack(link8, 1)
	PrintList(link8)
	ListRemoveIf(link8, 1)
	fmt.Println("------answer-----")
	PrintList(link8)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link7, 1)
	ListPushBack(link7, "Hello")
	ListPushBack(link7, 1)
	ListPushBack(link7, "There")
	ListPushBack(link7, 1)
	ListPushBack(link7, 1)
	ListPushBack(link7, "How")
	ListPushBack(link7, 1)
	ListPushBack(link7, "are")
	ListPushBack(link7, "you")
	ListPushBack(link7, 1)
	PrintList(link7)

	ListRemoveIf(link7, 1)
	fmt.Println("------answer-----")
	PrintList(link7)



	link9 := &List{}
	link10 := &List{}

	ListPushBack(link9, "a")
	ListPushBack(link9, "b")
	ListPushBack(link9, "c")
	ListPushBack(link9, "d")
	fmt.Println("-----first List------")
	PrintList(link9)

	ListPushBack(link10, "e")
	ListPushBack(link10, "f")
	ListPushBack(link10, "g")
	ListPushBack(link10, "h")
	fmt.Println("-----second List------")
	PrintList(link10)

	fmt.Println("-----Merged List-----")
	ListMerge(link9, link10)
	PrintList(link9)
}


