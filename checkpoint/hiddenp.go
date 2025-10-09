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
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return 
	}

	for i := 0 ; i < len(hidden) ;i++{
		for j := 0 ; j < len(main_sentence) ;j++{
			if hidden[i] == main_sentence[j]{
				if main_sentence[j : j+len(hidden)] == hidden{
					z01.PrintRune('1')
					z01.PrintRune('\n')
					return
				}
			}
		}
	}
	z01.PrintRune('0')
	z01.PrintRune('\n')

}