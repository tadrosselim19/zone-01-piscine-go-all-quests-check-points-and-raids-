package main

import(
	"github.com/01-edu/z01"
	"os"
)

func amain(){
	names := os.Args[1]
	slashespostion := []int{}
	for i := 0 ; i < len(names) ; i++{
		if names[i] == ' ' {
			slashespostion = append(slashespostion, i+1)
		}
	}
	for j := 0 ; j < len (slashespostion) ; j++{
		for _, a :=range (names[slashespostion[j]]){
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}

func main(){
	names := os.Args[1:]
	for _, i := range(names){
		for _, j := range(i){
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}