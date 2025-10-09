package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}
	pre_joined := os. Args[1:]
	arguments := ""
	// joining the argument 
	for i := 0 ; i < len(pre_joined) ; i++{
		for j := 0 ; j < len(pre_joined[i]) ; j++{
			arguments += string(rune(pre_joined[i][j]))
		}
		arguments += " "
	}
	vowels := "aAeEiIoOuU"
	reversed_vowels := ""
	reversed_vowels_inndex := []int{}
	// finding the vowels and stor 
	for i := 0 ; i < len(arguments) ; i++{
			for k := 0 ; k < len(vowels) ; k++{
				if arguments[i] == vowels[k]{
					reversed_vowels += string(rune(arguments[i]))
					reversed_vowels_inndex = append(reversed_vowels_inndex, i)
				}

			}
		}
	// printing and relace in the same time
	reversed_vowels_length := len(reversed_vowels)-1
	n := 0 
	    for i := 0 ; i < len(arguments) ; i++{
			if n < len(reversed_vowels_inndex) && i == reversed_vowels_inndex[n]{
				z01.PrintRune(rune(reversed_vowels[reversed_vowels_length]))
				reversed_vowels_length--
				n++
				continue
			}
			z01.PrintRune(rune(arguments[i]))
		}
	
	
}





