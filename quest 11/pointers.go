package main

import "fmt"

func main() {
	x := 5

	p1 := &x

	p2 := &p1

	p3 := &p2

	
	var &p4 int = x


	fmt.Println(*p3)
	fmt.Println(**p3)
	fmt.Println(***p3)


}