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