package main

import "github.com/01-edu/z01"

func PrintMemory1(arr [10]byte) {
	r := []string{}
	for _, byt := range arr {
		if byt == 0 {
			r = append(r, "00")
		} else {
			r = append(r, hex(int(byt)))
		}
	}

	// z01.PrintRune(rune(31)) 	// testing printable character min limits
	// z01.PrintRune(rune(127))	// testing printable character max limits

	// printing r:
	for i := 0; i < len(r); i++ {
		ppstr(r[i])
		z01.PrintRune(' ')
		if (i-3)%4 == 0 || i == len(r)-1 {
			z01.PrintRune('\n') // print new line each four strings and in the last one
		}
	}
	// printing ASCII graphic characters
	for _, byt := range arr {
		if byt < 32 || byt > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func ppstr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func hex(n int) string {
	base := "0123456789abcdef"
	sliceBase := []rune(base)
	index := []int{}
	var r []rune
	for n > 0 {
		index = append(index, n%16)
		n /= 16
	}
	for i := len(index) - 1; i >= 0; i-- {
		r = append(r, sliceBase[index[i]])
	}
	return string(r)
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// checkpoint / hardcoded using append (i don't know if it's allowed)

package main

import "github.com/01-edu/z01"

func PrintMemory2(arr [10]byte) {
	hex := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(hex[arr[i]>>4]))
		z01.PrintRune(rune(hex[arr[i]&15]))
		z01.PrintRune(' ')
		if (i-3)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		}
	}
	for _, byt := range arr {
		if byt < 32 || byt > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// checkpoint


/************************* Explanation***************************
these two lines saved us too much code:
z01.PrintRune(rune(hex[arr[i]>>4]))
z01.PrintRune(rune(hex[arr[i]&15]))

so let's try to understand them:

The conversion from binary to hexadecimal is efficient 
and straightforward because both are base-2 systems, 
meaning hexadecimal is a base-16 system which is a power 
of 2 (2^4 = 16). This relationship allows for a direct 
and simple conversion process, unlike the conversion 
between binary and decimal, where the base-10 system does 
not align as neatly with binary's base-2 system.

Each hexadecimal digit can represent four binary digits 
(bits) because 2^4 = 16. Thus, one hexadecimal digit can 
exactly represent a range from 0000 to 1111 in binary 
(or 0 to 15 in decimal).

Since one hexadecimal digit corresponds to exactly 4 bits 
in binary, we can split a byte (8 bits) into two halves, 
with each half directly mapping to a single hexadecimal 
digit. This is why the method involves isolating the 
high and low 4-bit parts of the byte.

High-order Half (arr[i]>>4):
By shifting the byte to the right by 4 bits (>>4), you 
effectively discard the lower 4 bits and move the upper 
4 bits to the position of the lower 4 bits. This 
operation transforms the upper half of the byte into 
a value that can be directly looked up in a hexadecimal 
mapping. For any given byte, this reveals the hexadecimal 
digit for the more significant part (the "high-order" half).

Low-order Half (arr[i]&15):
Applying a bitwise AND operation with 15 (0x0F or 
00001111 in binary) masks the upper 4 bits of the 
byte, isolating the lower 4 bits. This is because 
any bit ANDed with 0 results in 0, and any bit 
ANDed with 1 remains unchanged. This operation 
reveals the hexadecimal digit for the less significant 
part (the "low-order" half).

*/

package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	sliceB := []rune(base)
	var n int
	x := len(sliceB)
	var r string
	index := []int{}
	if len(sliceB) < 2 {
		printstr("NV")
		return
	}
	for i := 0; i < len(sliceB); i++ {
		for j := 0; j < len(sliceB); j++ {
			if i != j && sliceB[i] == sliceB[j] {
				printstr("NV")
				return
			}
		}
		if sliceB[i] == '+' || sliceB[i] == '-' {
			printstr("NV")
			return
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		n = -nbr
	} else if nbr > 0 {
		n = nbr
	} else {
		z01.PrintRune('0')
	}
	if nbr == -9223372036854775808 { // i have a problem for this particular case
		printstr("9223372036854775808")
	}
	for n > 0 {
		index = append(index, n%x)
		n /= x
	}
	for i := 0; i < len(index); i++ {
		r = string(sliceB[index[i]]) + r
	}
	printstr(r)
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}

// quest 5 / checkpoint

package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	for _, str := range args {
		for i := 0; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' && ((i+1 < len(str) && str[i+1] == ' ') || (i == len(str)-1)) { // the condition str[i] >= 'a' && str[i] <= 'z' is factorized
				os.Stdout.WriteString(string(str[i] - 32))
			} else if str[i] >= 'A' && str[i] <= 'Z' && i+1 < len(str) && str[i+1] != ' ' && i != len(str)-1 {
				os.Stdout.WriteString(string(str[i] + 32))
			} else {
				os.Stdout.WriteString(string(str[i]))
			}
		}
		os.Stdout.WriteString("\n")
	}
}

// checkpoint

package main

import "os"

// import "fmt"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	if args[0] == "" {
		os.Stdout.WriteString("\n")
		return
	}
	strSlice := splitt(args[0], " ")
	result := strSlice[len(strSlice)-1]
	for i := len(strSlice) - 2; i >= 0; i-- {
		result = result + " " + strSlice[i]
	}
	os.Stdout.WriteString(result + "\n")
}

func splitt(s, sp string) []string {
	strs := []string{}
	index := 0
	for i := 0; i+len(sp) <= len(s); i++ {
		if s[i:i+len(sp)] == sp {
			if s[index:i] != "" {
				strs = append(strs, s[index:i])
			}
			index = i + len(sp)
		} else if i+len(sp) == len(s) { // if len s[i:i+len(sp)]!= sep and i+len(sp)=len(s) (WHIICH MEANS we are at the last iteration) WHIICH MEANS if len s[len(s)-len(sep):len(s))]!= sep WHIICH MEANS if the end of the string doesn't end with a sep
			strs = append(strs, s[index:]) // in this case append all to the end
		}
	}
	// fmt.Printf("%#v\n", strs) // this line allow us to see clearly how and why we added the condition "if s[index:i]!=""{" above
	return strs
}


// checkpoint