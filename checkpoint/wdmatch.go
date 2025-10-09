package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args)!= 3 || len(os.Args[1]) > len(os.Args[2]) {
		return
	}

	target := os.Args[1]
	search := os.Args[2]
	start := 0
	found := []bool{}
// check every charachter presence in order
	for i := 0 ; i < len(target) ; i++{
        for j := start ; j < len(search) ; j++{
			if target[i] == search[j] {
				found = append(found, true)
				start = j+1
				break
			}
	    }
	}
// printing the result
if len(found) == len(target){
	for _, ch := range(target){
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}else{
	return
}

}