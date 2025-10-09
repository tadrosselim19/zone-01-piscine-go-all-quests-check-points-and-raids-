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
	count := 0
	check_value :=""
	stared := 0
	for i := 0 ; i < len(hidden) ;i++{
		for j := stared ; j < len(main_sentence) ;j++{
			if hidden[i] == main_sentence[j] {
				count++
				check_value += string(hidden[i])
				stared = j+1
				break
			}
		}
	}
	if len(hidden) == count && check_value == hidden{
        z01.PrintRune('1')
	    z01.PrintRune('\n')
	}else{
       z01.PrintRune('0')
	   z01.PrintRune('\n')
	}
}