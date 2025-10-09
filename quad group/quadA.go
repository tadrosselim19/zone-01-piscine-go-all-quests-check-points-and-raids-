package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for col := 1; col <= y; col++ {
		for row := 1; row <= x; row++ {
			if (col == 1) || (col == y) {
				if (row == 1) || (row == x) {
					fmt.Print('o')
				} else {
					fmt.Print('-')
				}
			} else {
				if (row == 1) || (row == x) {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}

}
func main() {
	x,_ := strconv.Atoi(os.Args[1])
	y,_ := strconv.Atoi(os.Args[2])
	QuadA(x,y)
}
