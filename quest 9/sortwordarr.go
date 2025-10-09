package main

import (
	"fmt"
)

func SortWordArr(a []string) {
	for i := 0 ; i < len(a) ; i++{
		for j := 0 ; j <len(a) ; j++{
			if a[i] < a[j]{
				a[i] , a[j] = a[j] , a[i]
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}