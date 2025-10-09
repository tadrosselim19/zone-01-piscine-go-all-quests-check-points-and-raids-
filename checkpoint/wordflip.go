package main

import (
	"fmt"
	"sort"
)
func WordFlip(str string) string {
	final := ""
	str_rune := []rune(str)
	reversed_array:= []string{}
	word := ""
    // sclicing the string
	
	flag := false
	for i := 0 ; i < len(str_rune); i++{
		if flag == false && (str_rune[i] == ' ' || str_rune[i] == '\t'){
			continue
		}
		if str_rune[i] != ' ' && str_rune[i] != '\t'{
			flag = true
			word += string(str_rune[i])
		}else{
			if word != ""{
		reversed_array = append([]string{word},reversed_array...)
	     }
			word = ""
		}
	}
	if word != ""{
		reversed_array = append([]string{word},reversed_array...)
	}
	// finalized the final string 
	for i := 0 ; i < len(reversed_array) ; i++{
		for j := 0 ; j < len(reversed_array[i]) ; j++{
			final += string(rune(reversed_array[i][j]))
		}
		if i != len(reversed_array)-1{
		final += " "
		}
	}

return final +"\n"
}


func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}