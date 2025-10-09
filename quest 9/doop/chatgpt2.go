package main

import (
	"os"
)

func writeInt(n int64) {
	buf := []byte{}
	if n == 0 {
		os.Stdout.Write([]byte{'0', '\n'})
		return
	}
	if n < 0 {
		buf = append(buf, '-')
		n = -n
	}
	digits := []byte{}
	for n > 0 {
		d := n % 10
		digits = append([]byte{byte(d + '0')}, digits...)
		n /= 10
	}
	buf = append(buf, digits...)
	buf = append(buf, '\n')
	os.Stdout.Write(buf)
}

func writeString(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	// --- Parse first number ---
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

	// --- Parse second number ---
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

	// --- Operator ---
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
		writeInt(res)
	case "-":
		res := value1 - value2
		if (value1 > 0 && value2 < 0 && res < 0) || (value1 < 0 && value2 > 0 && res > 0) {
			return
		}
		writeInt(res)
	case "*":
		res := value1 * value2
		if value2 != 0 && res/value2 != value1 {
			return
		}
		writeInt(res)
	case "/":
		if value2 == 0 {
			writeString("No division by 0")
			return
		}
		writeInt(value1 / value2)
	case "%":
		if value2 == 0 {
			writeString("No modulo by 0")
			return
		}
		writeInt(value1 % value2)
	default:
		return
	}
}
