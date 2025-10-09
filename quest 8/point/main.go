package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 6 * 7 // 42
	ptr.y = 3 * 7 // 21
}

func main() {
	output := []rune{
		120, // 'x'
		32,  // ' '
		61,  // '='
		32,  // ' '
		52,  // '4'
		50,  // '2'
		44,  // ','
		32,  // ' '
		121, // 'y'
		32,  // ' '
		61,  // '='
		32,  // ' '
		50,  // '2'
		49,  // '1'
		10,  // '\n'
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ten := 'B' - '@' // 2
	one := 'B' - 'A' // 1

	ptr.x = int(ten*ten*ten*ten + ten*ten) // 2^4 + 2^2 = 16 + 4 = 20 (adjust as needed)
	ptr.y = int(ten*ten + one)             // 4 + 1 = 5 (adjust as needed)
}

func main() {
	points := &point{}
	setPoint(points)

	ten := 'B' - '@' // 2
	one := 'B' - 'A' // 1
	two := ten
	four := two + two
	zero := 'A' - ('A' - '0')

	output := []rune{
		'x', ' ', '=', ' ',
		zero + four,
		zero + two,
		',', ' ',
		'y', ' ', '=', ' ',
		zero + two,
		zero + one,
		'\n',
	}

	for _, r := range output {
		z01.PrintRune(r)
	}
}
package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ten := 'B' - '1' // = 17
	two := '4' - '2' // = 2
	one := 'A' - '@' // = 1

	ptr.x = int(two*ten + two) // 2*10 + 2 = 22
	ptr.x += int(two * ten)    // 22 + 20 = 42

	ptr.y = int(two*ten + one) // 2*10 + 1 = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")
	printNbr(points.x)
	printStr(", y = ")
	printNbr(points.y)
	z01.PrintRune('\n')
}
