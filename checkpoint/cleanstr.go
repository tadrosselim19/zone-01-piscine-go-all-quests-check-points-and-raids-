package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	sentence := os.Args[1]
	if sentence == ""{
		fmt.Println()
		return
	}
	pre_final :=""
	final := ""
	// loop for remving multi spaces next each other 
	i := 0
		for j := i+1 ; j< len(sentence);j++{
			if !((sentence[i]== ' ' && sentence [j] == ' ') || (sentence[i]== '\t' && sentence [j] == '\t')){
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
	fmt.Println(final)
	
}