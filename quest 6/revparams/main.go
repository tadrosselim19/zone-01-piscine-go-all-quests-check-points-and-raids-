package main

import(
	"github.com/01-edu/z01"

	"os"
)

func main(){
	names := os.Args[1:]

	for i:= len(names)-1 ; i >= 0 ;i--{
		for _, j := range(names[i]){
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}