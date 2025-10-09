package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4{
		return
	}
	string_main := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]
	final := ""
	for i := 0 ; i < len(string_main) ;{
		if i+len(search)-1 < len(string_main) && string_main[i : i+len(search)] == search{
			final += replace
			i += len(search)
		}else{
			final+= string(rune(string_main[i]))
			i++
		}
	}
	for _, ch := range(final) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}