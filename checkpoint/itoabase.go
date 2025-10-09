package main

import "fmt"

func ItoaBase(value, base int) string {
	base_base := "0123456789ABCDEF"
	final := ""
	index_array := []int{}
	if value == 0{
		return "0"
	}
	if value *-1 > 0{
		value *= -1
		final+="-" 
	}
	for value >= 1{
		index := value % base
		index_array = append([]int{index},index_array...)
		value = value / base
	}
	for i := 0 ; i < len(index_array) ; i++{
		final += string(rune(base_base[index_array[i]]))
	}
	return final

}
func main() {
	fmt.Println(ItoaBase(10, 2))     // "1010"
	fmt.Println(ItoaBase(255, 16)) // "FF"
	fmt.Println(ItoaBase(-13, 4))  // "-31"
}