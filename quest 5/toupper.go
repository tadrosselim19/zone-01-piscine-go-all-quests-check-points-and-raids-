package main

import (
	"fmt"
)

func aToUpper(s string) string {
	srunes :=[]rune(s)
	result := ""
	for _, i := range(srunes){
			if i >= 'A' && i <= 'Z'{
				result += string(i)
			}else if i >= 'a' && i <= 'z'{
				result += string(i-32)
			}else{
				result += string(i)
			}
	}
	return result
}

func ToUpper(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(r - 32)
		} else {
			result += string(r)
		}
	}
	return result
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}