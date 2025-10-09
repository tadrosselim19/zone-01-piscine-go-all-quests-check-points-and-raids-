package main

import (
	"fmt"
	"os"
)

func validation_of_target_and_sclise(s string)([]string,bool){
	final := []string{}
	word := ""
	for i := 0 ; i < len(s) ; i++{
		if s[0] != '(' || s[len(s)-1] != ')'{
			return nil ,false
		}
		if i == 0 && s[i] == '(' {
			continue
		}
		if i == len(s)-1 &&  s[i] == ')'{
			final = append(final, word)
			word = ""
		}
		if s[i] != '|'{
			word += string(rune(s[i]))
		}else{
			final = append(final, word)
			word = ""
		}
	}
	return final , true
}
func sclising(s string)[]string{
	final := []string{}
	word := ""
	for i := 0 ; i < len(s) ; i++{
		if s[i] == ',' || s[i] == '.'{ 
			continue
		}
		if s[i] != ' ' {
			word+= string(rune(s[i]))
		}else{
			final = append(final, word)
			word = ""
		}
	}
	if word != ""{
		final = append(final, word)
	}
	return final
}
func main() {
	if len(os.Args) != 3 {
		return
	}
	if os.Args[1] == "" || os.Args[2] == "" {
		return
	}
	target := os.Args[1]
	new_target , flag := validation_of_target_and_sclise(target)
	if flag == false{
		return
	}
	serching := os.Args[2]
	scliced_seaching := sclising(serching)
	// loop for searching and printing 
	count := 0
	for j := 0 ; j < len(scliced_seaching) ; j++{
		for i := 0 ; i < len(new_target) ; i++ {
		   for k := 0 ; k < len(scliced_seaching[j]) ; k++{
			  if  k + len(new_target[i])-1 < len(scliced_seaching[j]) && new_target[i] == scliced_seaching[j][k:k+len(new_target[i])]{
				count++
				fmt.Printf("%v: %v\n",count,scliced_seaching[j])
				break
				
			}
		}	
		}
	}

}