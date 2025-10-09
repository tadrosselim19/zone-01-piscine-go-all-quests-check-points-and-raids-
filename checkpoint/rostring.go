package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2{
		z01.PrintRune('\n')
		return
	}
	arguments := os.Args[1:]
	array_of_all_arguments := [][]string{}

	for i := 0 ; i < len(arguments) ; i++{
		array_of_words := []string{}
		words := ""

		stated := 0
		for j := 0 ; j < len(arguments[i]) ; j++{
			if (arguments[i][j] == ' ' || arguments[i][j] == '\t') && stated == 0 {
				continue
			}
			stated = -1
			if j+1 != len(arguments[i]) && (arguments[i][j] == ' ' || arguments[i][j] == '\t') && (arguments[i][j+1] == ' ' || arguments[i][j+1] == '\t') {
				continue
			}
			if j == len(arguments[i])-1 && (arguments[i][j] == ' ' || arguments[i][j] == '\t'){
				break
			}
			if arguments[i][j] == ' ' || arguments[i][j] == '\t'{
				array_of_words = append(array_of_words, words)
				words = ""
			}else{
			words += string(rune(arguments[i][j]))
			}
		}
		array_of_words = append(array_of_words, words)
		array_of_all_arguments =append(array_of_all_arguments, array_of_words)
	}
	fmt.Println(array_of_all_arguments)
	//printing and relacing

	for i:=0 ; i < len(array_of_all_arguments) ; i++{
		
		for j:=1 ; j < len(array_of_all_arguments[i]) ; j++{
			for k:= 0 ; k < len(array_of_all_arguments[i][j]) ; k++{
					z01.PrintRune(rune(array_of_all_arguments[i][j][k]))
			}
			z01.PrintRune(' ')
		}

		// printing the frist word

		for j:=0 ; j < 1 ; j++{
			for k:= 0 ; k < len(array_of_all_arguments[i][j]) ; k++{
					z01.PrintRune(rune(array_of_all_arguments[i][j][k]))
			}
		}
		z01.PrintRune('\n')
	}
}
