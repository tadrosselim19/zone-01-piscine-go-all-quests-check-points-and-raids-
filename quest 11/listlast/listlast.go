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
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))

}
