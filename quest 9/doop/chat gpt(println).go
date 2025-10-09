package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	// --- Parse first value ---
	s1 := os.Args[1]
	var value1 int64
	sign1 := int64(1)
	i := 0
	if s1[0] == '-' {
		sign1 = -1
		i++
	} else if s1[0] == '+' {
		i++
	}
	if i == len(s1) {
		return
	}
	for ; i < len(s1); i++ {
		if s1[i] < '0' || s1[i] > '9' {
			return
		}
		d := int64(s1[i] - '0')
		if value1 > (9223372036854775807-d)/10 {
			return
		}
		value1 = value1*10 + d
	}
	value1 *= sign1

	// --- Parse second value ---
	s2 := os.Args[3]
	var value2 int64
	sign2 := int64(1)
	j := 0
	if s2[0] == '-' {
		sign2 = -1
		j++
	} else if s2[0] == '+' {
		j++
	}
	if j == len(s2) {
		return
	}
	for ; j < len(s2); j++ {
		if s2[j] < '0' || s2[j] > '9' {
			return
		}
		d := int64(s2[j] - '0')
		if value2 > (9223372036854775807-d)/10 {
			return
		}
		value2 = value2*10 + d
	}
	value2 *= sign2

	// --- Operator check ---
	op := os.Args[2]
	if len(op) != 1 {
		return
	}

	switch op {
	case "+":
		res := value1 + value2
		if (value1 > 0 && value2 > 0 && res < 0) || (value1 < 0 && value2 < 0 && res > 0) {
			return
		}
		println(res)
	case "-":
		res := value1 - value2
		if (value1 > 0 && value2 < 0 && res < 0) || (value1 < 0 && value2 > 0 && res > 0) {
			return
		}
		println(res)
	case "*":
		res := value1 * value2
		if value2 != 0 && res/value2 != value1 {
			return
		}
		println(res)
	case "/":
		if value2 == 0 {
			println("No division by 0")
			return
		}
		println(value1 / value2)
	case "%":
		if value2 == 0 {
			println("No modulo by 0")
			return
		}
		println(value1 % value2)
	default:
		return
	}
}
