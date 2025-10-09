package main

import (
	"fmt"
	"os"
)
func isWhitespace(c byte) bool {
    return c == ' ' || c == '\t'
}
func main() {
	if len(os.Args) != 2 {
		
		return
	}
	sentence := os.Args[1]
	if sentence == "" {
		
		return
	}
	pre_final :=""
	final := ""
	post_final:= ""
	// loop for remving multi spaces next each other 
	i := 0
		for j := i+1 ; j< len(sentence);j++{
			if !(isWhitespace(sentence[i]) && isWhitespace(sentence[j])) {
             pre_final += string(sentence[i])
            }
			i++
		}
	pre_final += string(sentence[i])

	// loop for remove behind and front spaces 
	for j := 0 ; j< len(pre_final);j++{
			if (pre_final[j] == ' ' && j == len(pre_final)-1) || (pre_final[j] == '\t' && j == len(pre_final)-1){
				break
			}
			if (pre_final[j] == ' ' && j == 0) || (pre_final[j] == '\t' && j == 0 ){
				continue
			}else{
				final += string(pre_final[j])
			}
		}
	
	// adding 3 spaces between each word

	for  j := 0 ; j< len(final);j++{
		if final[j] != ' '{
			post_final += string(final[j])
		}else{
			post_final += "   "
		}		
	}
	fmt.Println(post_final)
}