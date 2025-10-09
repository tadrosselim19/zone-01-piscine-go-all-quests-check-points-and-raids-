package main

import "fmt"


func main() {
	size := 10 
	fmt.Println(Created(size))
}


func Created(size int)[]int {
	s := make([]int , size)
	for i:=1 ; i <= size ; i++{
		s[i-1]=i
	}
	return s

}