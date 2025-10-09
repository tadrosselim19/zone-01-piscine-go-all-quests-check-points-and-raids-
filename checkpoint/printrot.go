package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) !=2 ||len(os.Args[1]) > 1{
		z01.PrintRune('\n')
		return
	}
	target := os.Args[1]
	end := 26
	flag := 0
	for i := 0 ; i < end ; i++{
		if byte(i + 'a') == target[0] && flag == 0{
			for j := target[0] ; j <= 'z' ;j++{
				z01.PrintRune(rune(j))
			}
			flag = 1
		}
		if flag == 1{
		   i=0
		   end = int(target[0]- 'a')
		   flag = 2
		}
		if flag == 2{
		   z01.PrintRune(rune(i+'a'))
		}

	}

}
func validad(){
	if os.Args[1][0] < 'a' || os.Args[1][0] > 'z' {
    z01.PrintRune('\n')
    return
	target := os.Args[1]
	alphabet := "abcdefghijklmnopqrstuvwxyz"
    start := int(target[0] - 'a')
    rotated := alphabet[start:] + alphabet[:start]
    for _, r := range rotated {
    z01.PrintRune(r)
    }
     z01.PrintRune('\n')

}

}