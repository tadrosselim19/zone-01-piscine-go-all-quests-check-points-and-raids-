package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arguments := os.Args[1:]
	len_argument := len(arguments)

	// convert all leters to small
	small_arguments := []string{}
	sentence := ""

	for i := 0 ; i < len_argument ; i++{
		for j := 0 ; j < len(arguments[i]) ; j++{
			if arguments[i][j] >= 'A' && arguments[i][j] <= 'Z'{
				sentence += string(rune(arguments[i][j] + 32))
			}else{
				sentence += string(rune(arguments[i][j]))
			}
		}
		small_arguments = append(small_arguments, sentence)
		sentence = ""
	}

	
	// convert the last letter to capital
	for i := 0 ; i < len(small_arguments) ; i++{
		for j := 0 ; j < len(small_arguments[i]) ; j++{
			if  j+1 != len(small_arguments[i]) && small_arguments[i][j] >= 'a' && small_arguments[i][j] <= 'z' && small_arguments[i][j+1] == ' ' {
				z01.PrintRune(rune(small_arguments[i][j] - 32))
			}else if j == len(small_arguments[i])-1 && small_arguments[i][j] >= 'a' && small_arguments[i][j] <= 'z'{
				z01.PrintRune(rune(small_arguments[i][j] - 32))
			}else{
				z01.PrintRune(rune(small_arguments[i][j]))
			}
		}
		z01.PrintRune('\n')
	}
}