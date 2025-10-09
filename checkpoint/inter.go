package main

import (
	
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	hidden  := os.Args[1]
	main_sentence := os.Args[2]

	if hidden == ""{
		z01.PrintRune('\n')
		return 
	}
	pre_final := ""
	for i := 0 ; i < len(hidden) ;i++{
		for j := 0 ; j < len(main_sentence) ;j++{
			if hidden[i] == main_sentence[j] {
				pre_final += string(hidden[i])
				break
			}
		}
	}

	for i := 0 ; i < len(pre_final);i++{
		found := false
		for j := 0 ; j < i ;j++{
			if pre_final[i] == pre_final[j] {
				found = true
				continue
			}
		}
		if found ==false {
				z01.PrintRune(rune(pre_final[i]))
			}
		
	}
	z01.PrintRune('\n')
}