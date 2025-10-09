package main

import "github.com/01-edu/z01"

func assci_num(arr [10]byte) []int {
	final := []int{}
	for i := 0; i < len(arr); i++ {
		final = append(final, int(arr[i]))
	}
	return final
}

func hex_conversion(arr []int) []string {
	base_hex := "0123456789abcdef"

	final := []string{}
	word := ""
	for i := 0; i < len(arr); i++ {
		index_arry := []int{}
		x := 1
		if arr[i] < 16 {
			index_arry = append(index_arry, 0) // ensure leading zero
		}
		if arr[i] == 16{
			word += "10"
			final = append(final, word)
			word =""
			continue
		}
		for j := 1; j < arr[i]; j *= 16 {
			x = j
		}
		for x >= 1 {
			index := arr[i] / x
			index_arry = append(index_arry, index)
			arr[i] = arr[i] % x
			x = x / 16
		}
		// loop for adding every ch of num at it index
		for k := 0; k < len(index_arry); k++ {
			word += string(base_hex[index_arry[k]])
		}
		final = append(final, word)
		word = ""
	}
	return final
}
func PrintMemory(arr [10]byte) {
	test_arry := assci_num(arr)
	final_arry := hex_conversion(test_arry)
	count := 0
	for i := 0; i < len(final_arry); i++ {
		for j := 0; j < len(final_arry[i]); j++ {
			z01.PrintRune(rune(final_arry[i][j]))
		}
		count++
		if count == 4{
			z01.PrintRune('\n')
			count = 0
		}else{
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	// printing the word
	for i := 0 ; i < len(arr) ; i++{
		if arr[i] >= 32 && arr[i] <= 126{
		   z01.PrintRune(rune(arr[i]))
		}else{
			z01.PrintRune('.')
		}
	}


}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

package main

import "github.com/01-edu/z01"

func asciiNums(arr [10]byte) []int {
	result := []int{}
	for _, b := range arr {
		result = append(result, int(b))
	}
	return result
}

func hexConversion(arr []int) []string {
	baseHex := "0123456789abcdef"
	result := []string{}
	for _, v := range arr {
		high := v / 16
		low := v % 16
		result = append(result, string(baseHex[high])+string(baseHex[low]))
	}
	return result
}

func PrintMemory(arr [10]byte) {
	hexArr := hexConversion(asciiNums(arr))

	// Print hex values with 4, 4, remainder grouping
	for i := 0; i < len(hexArr); i++ {
		for _, ch := range hexArr[i] {
			z01.PrintRune(ch)
		}
		if (i+1)%4 == 0 || i == len(hexArr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	// Print ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
