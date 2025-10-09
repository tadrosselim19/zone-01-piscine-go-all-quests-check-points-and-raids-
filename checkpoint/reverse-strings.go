package main

import (
	"fmt"

	
)
func ReverseStrings(strs []string) string {
	final :=""
	for i := len(strs)-1 ; i >= 0  ; i--{
		for j := len(strs[i])-1; j >=0 ;j--{
			final += string(rune(strs[i][j]))
		}
		if i != 0{
			final+= " "
		}
		
	}
	return final

}

func main(){
    fmt.Println(ReverseStrings([]string{"a", "b", "c"}))
    fmt.Println(ReverseStrings([]string{"Good","Morning!"}))
    fmt.Println(ReverseStrings([]string{"Hello World"}))
}