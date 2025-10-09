package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	argsInt := []int{}
	result := 0

	if len(args) == 0 {
		return
	}

	if args[0] == "--upper" {
		upper = true
	}

	if upper == false {
		for i := 0; i < len(args); i++ {
			for j := 0; j < len(args[i]); j++ {
				if args[i][j] <= '9' && args[i][j] >= '0' {
					result = result*10 + int(args[i][j]-'0')
				} else {
					argsInt = append(argsInt, 0)
				}
			}
			argsInt = append(argsInt, result)
			result = 0
		}
	}

	if upper == true {
		for i := 1; i < len(args); i++ {
			for j := 0; j < len(args[i]); j++ {
				if args[i][j] <= '9' && args[i][j] >= '0' {
					result = result*10 + int(args[i][j]-'0')
				} else {
					argsInt = append(argsInt, 0)
				}
			}
			argsInt = append(argsInt, result)
			result = 0
		}
	}

	for k := 0; k < len(argsInt); k++ {
		if upper == false {
			if k != len(argsInt)-1 {
			if argsInt[k] <= 26 && argsInt[k] >= 1 {
				z01.PrintRune(rune(argsInt[k] + 96))
			} else {
				z01.PrintRune(' ')
			}
		}else {
				if argsInt[k] <= 26 && argsInt[k] >= 1 {
					z01.PrintRune(rune(argsInt[k] + 96))
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(' ')
					z01.PrintRune('\n')
				}
			}
		}
		if upper == true {
			if k != len(argsInt)-1 {
				if argsInt[k] <= 26 && argsInt[k] >= 1 {
					z01.PrintRune(rune(argsInt[k] + 64))
				} else {
					z01.PrintRune(' ')
				}
			} else {
				if argsInt[k] <= 26 && argsInt[k] >= 1 {
					z01.PrintRune(rune(argsInt[k] + 64))
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(' ')
					z01.PrintRune('\n')
				}
			}
		}
	}
}
