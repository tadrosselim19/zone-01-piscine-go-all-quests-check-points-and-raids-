package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){
	name := os.Args[0]
	lastslah := 0
	for i := len(name)-1 ; i >=0 ; i--{
		if name[i] == '/'{
			lastslah = i
			break
		}
	}
	programname:=name[lastslah+1:]

	for _,j := range(programname){
		z01.PrintRune(j)

	}

	
	z01.PrintRune('\n')
	 
}
func amain() {
	name := os.Args[0]
	lastSlash := 0

	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			lastSlash = i
			break
		}
	}

	// Get the substring after the last slash
	progName := name[lastSlash+1:]

	// Print each rune properly
	for _, r := range progName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
