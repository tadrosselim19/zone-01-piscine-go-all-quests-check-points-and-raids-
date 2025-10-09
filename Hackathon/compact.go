package main

import (
	"fmt"
)

const N = 6
func Compact(ptr *[]string) int {
	sclice := *ptr
	result := []string{}
	for i := 0 ; i < len(sclice) ; i++{
		if sclice[i] != ""{
			result = append(result, sclice[i])
		}
	}
	*ptr = result
	return len(result)

}

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}