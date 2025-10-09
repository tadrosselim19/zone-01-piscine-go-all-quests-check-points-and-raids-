package main

import "fmt"

type node struct{
	data int
	next *node
}

type linklist struct{
	head *node
	length int
}

func (l *linklist) prepend(n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length ++ 
}

func (l linklist) aprintlistdata(){
	toprint := l.head
	for l.length != 0{
		fmt.Printf("%d ->", toprint.data)
		toprint = toprint.next
		l.length--
	}
	fmt.Println(nil)
}
func (l linklist) printlistdata() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println(nil)
}

func (l *linklist) deletevalue(value int) {
	if l.length == 0 {
		return
	}
	// If the head node is the one to delete
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	// Traverse the list to find the node before the one to delete
	current := l.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			l.length--
			return
		}
		current = current.next
	}
}


func main(){
	mylist := linklist{}
	node1:= &node{data: 1}
	node2:= &node{data: 2}
	node3:= &node{data: 3}
	node4:= &node{data: 4}
	node5:= &node{data: 20}
	node6:= &node{data: 30}
	node7:= &node{data: 54}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)


	mylist.printlistdata()
	println()

	mylist.deletevalue(100)
	mylist.printlistdata()
	println()

	mylist.deletevalue(4)
	mylist.printlistdata()
	println()
	
	mylist.deletevalue(1)
	mylist.printlistdata()
	println()

	emptylist := linklist{}
	emptylist.printlistdata()
	println()


}

