package main

import (
	"fmt"
	
)

func AlphaCount(s string) int {
	srunes := []rune(s)
	result := ""
	for _, i:=range(srunes){
		for j := 'a' ; j <= 'z' ; j++{
			if i == j {
				result +=string(j)
			}
		}
	}
	for _, l:=range(srunes){
		for k := 'A' ; k <= 'Z' ; k++{
			if l == k {
				result +=string(k)
			}
		}
	}
	return len(result)

}

func aAlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}
	return count
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}