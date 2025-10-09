package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for col := 1; col <= y; col++ {
		for row := 1; row <= x; row++ {
			if col == 1 {
				if row == 1 {
					z01.PrintRune('/')
				} else if row == x {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else if col == y {
				if row == 1 {
					z01.PrintRune('\\')
				} else if row == x {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else {
				if (row == 1) || (row == x) {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadB(x, y)
}