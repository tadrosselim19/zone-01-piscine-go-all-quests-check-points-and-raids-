package main

import(
	"github.com/01-edu/z01"

	"os"
)

func main(){
	names := os.Args[1:]

	
for l := 0; l < len(names)-1; l++ {
	for i := 0 ; i < len(names)-1 ; i++{
		if names[i] >  names[i+1]{
			names[i] ,names[i+1] = names[i+1] , names[i]
		}
	}
}
	for _, k := range(names){
		for _, j := range(k){
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}

func amain(){
	names := os.Args[1:]
	for i := 0 ; i < len(names) ; i++{
		if names[i] > names[i+1]{
			names[i] , names[i+1] =names[i+1] , names[i]
		}
	}
	for _, k := range(names){
		for _, j := range(k){
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}