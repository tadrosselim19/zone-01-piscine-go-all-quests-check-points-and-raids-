package main

import (
	"os"

	"github.com/01-edu/z01"
)


func convert4(num string)int{
	rune_num := []rune(num)
	final := 0	
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return -1
		}
	}
return final
}
func main() {
	
	if len(os.Args)!= 2 {
		return
	}
	pre_number := os.Args[1]
	number := convert4(pre_number)
	if number == -1 || number == 0{
		return
	}
	x := number
	flag := false
	for x > 0{
		if x == 1{
			flag = true
			break
		}else if x % 2 == 0{
			x = x/2
		}else {
			break
		}
	}

	if flag{
		z01.PrintRune('t')
		z01.PrintRune('r')
		z01.PrintRune('u')
		z01.PrintRune('e')
		z01.PrintRune('\n')

	}else{
		z01.PrintRune('f')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('s')
		z01.PrintRune('e')
		z01.PrintRune('\n')
	}


}