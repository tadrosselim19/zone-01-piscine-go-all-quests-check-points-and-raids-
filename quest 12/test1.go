package main

import (
	"fmt"
)
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}


func aBTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		root = newNode
		return root
	}

	current := root
	for {
		if data < current.Data { 
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				break
			}
			current = current.Right
		}
	}

	return root
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		root = newNode
		return root
	}

	current := root
	var Parent *TreeNode

	for current != nil {
		Parent = current
		if data <= current.Data { 
			current = current.Left
		} else  {
			current = current.Right
		}
	}

	if Parent.Data >= data{
		Parent.Left=newNode
	}else{
		Parent.Right=newNode
	}
	return root
}
func iBTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		root = newNode
		return root
	}
	return root
}
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.Right, f)
	f(root.Data)
	BTreeApplyInorder(root.Left, f)
	
}
func cBTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
    BTreeApplyInorder(root.Left, f)
	BTreeApplyInorder(root.Right, f)
	
	
	f(root.Data)
}
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)	
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Data == elem{
		return root
	}
	
	curent := root

	for curent != nil && curent.Data != elem{
		if elem > curent.Data{
			curent = curent.Right
		}else{
			curent = curent.Left
		}
	}
	
	return curent
}

func iBTreeLevelCount(root *TreeNode) int {
	if root == nil{
		return 0
	}
	currentrh := root
	currentlf := root
	countrh := 0
	countlf := 0
	for currentrh != nil {
		countrh++
		currentrh = currentrh.Right
	}
	for currentlf != nil {
		countlf++
		currentlf = currentlf.Left

	}
	if countrh > countlf{
		return countrh +1
	}
	return  countlf + 1
}
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}	


func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

func isBST(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Data <= *min {
		return false
	}
	if max != nil && node.Data >= *max {
		return false
	}

	return isBST(node.Left, min, &node.Data) && isBST(node.Right, &node.Data, max)
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		f(node.Data)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}



func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)
	println()
	BTreeApplyInorder(root, fmt.Println)
	println()
	BTreeApplyPostorder(root, fmt.Println)
	println()
	cBTreeApplyPostorder(root, fmt.Println)
	println()
	BTreeApplyPreorder(root, fmt.Println)

	selected := BTreeSearchItem(root, "7")
	fmt.Print("Item selected -> ")
	if selected != nil {
		fmt.Println(selected.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Parent of selected item -> ")
	if selected.Parent != nil {
		fmt.Println(selected.Parent.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Left child of selected item -> ")
	if selected.Left != nil {
		fmt.Println(selected.Left.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Right child of selected item -> ")
	if selected.Right != nil {
		fmt.Println(selected.Right.Data)
	} else {
		fmt.Println("nil")
	}
	println()
	fmt.Println(BTreeLevelCount(root))


	println()
	fmt.Println(BTreeIsBinary(root))   


    println()  
	BTreeApplyByLevel(root, fmt.Println) 

    println()
 	max := BTreeMax(root)
	fmt.Println(max.Data)

    println()
	min := BTreeMin(root)
	fmt.Println(min.Data)

	node := BTreeSearchItem(root, "1")
	rplc := &TreeNode{Data: "3"}
	root = BTreeTransplant(root, node, rplc)
	BTreeApplyInorder(root, fmt.Println)

}