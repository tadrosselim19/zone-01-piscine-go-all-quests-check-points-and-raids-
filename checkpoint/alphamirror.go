package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args)!=2{
		z01.PrintRune('\n')
		return
	}
	input:=os.Args[1]
	lo := "abcdefghijklmnopqrstuvwxyz"
	up := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lo_reversed := "zyxwvutsrqponmlkjihgfedcba"
	upper_reversed := "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	final:= ""

	for i:=0 ; i < len(input) ; i++{
		if input[i] >= 'a' && input[i] <= 'z'{
			for j:=0 ; j < len(lo) ; j++{
			if input[i] == lo[j]{
				final += string(lo_reversed[j])
			}
		}
		}else if input[i] >= 'A' && input[i] <= 'Z'{
			for j:=0 ; j < len(up) ; j++{
			if input[i] == up[j]{
				final += string(upper_reversed[j])
			}
		}
		}else{
			final += string(input[i])
		}
	}

	for i:= 0 ; i < len(final) ;i++{
		z01.PrintRune(rune(final[i]))
	}
	z01.PrintRune('\n')
	
}

// func main() {
// 	if len(os.Args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	input := os.Args[1]

// 	for _, c := range input {
// 		if c >= 'a' && c <= 'z' {
// 			z01.PrintRune('z' - (c - 'a'))
// 		} else if c >= 'A' && c <= 'Z' {
// 			z01.PrintRune('Z' - (c - 'A'))
// 		} else {
// 			z01.PrintRune(c)
// 		}
// 	}

// 	z01.PrintRune('\n')
// }