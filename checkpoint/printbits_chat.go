package main

import (
	"fmt"
	"os"
)
func printBits1(num int) {
	bits := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		bits[i] = byte(num % 2) + '0' // convert 0 or 1 to '0' or '1'
		num = num / 2
	}
	os.Stdout.Write(bits)
}
func main() {
	if len(os.Args) != 2 {
		return
	}

	var num int
	_, err := fmt.Sscanf(os.Args[1], "%d", &num)
	if err != nil {
		fmt.Println("invalid input")
		return
	}

	// Print 32-bit binary
	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1 // shift right and mask last bit
		fmt.Print(bit)
		if i%8 == 0 && i != 0 { // space after every byte
			fmt.Print(" ")
		}
	}

	
}
