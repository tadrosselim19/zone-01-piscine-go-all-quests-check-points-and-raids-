package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	
	vowels := "aeiouAEIOU"
	sentence := os.Args[1]

	stared := -1

	// loop for find the stareding point of index and print if find vowels
	for i := 0 ; i < len(sentence) ; i++{
		for j := 0 ; j < len(vowels) ; j++ {
			if vowels[j] == sentence[i]{
				stared = i
				for k := stared ; k < len(sentence) ; k++ {
		            z01.PrintRune(rune(sentence[k]))
	            }
	            
				for l := 0 ; l < stared ; l++{
					z01.PrintRune(rune(sentence[l]))
				} 
				z01.PrintRune('a')
	            z01.PrintRune('y')
	            z01.PrintRune('\n')
	            return 
			}
		}
	}
	// the program cant find any of vowels
		os.Stdout.WriteString("No vowels\n")
		return
	
}