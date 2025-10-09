package main

import (
	"fmt"

)
func power_math(len_base int, index int)int{
	mat_power := 1
	for i := 0 ; i <index ; i++{
		mat_power *= len_base
	}
	return mat_power
}
func convert_to_decimal(nbr string , baseFrom string)int{
	final := 0
	for i:=0 ; i< len(nbr) ;i++{
		for j:=0 ; j < len(baseFrom) ; j++{
			if baseFrom[j] == nbr[i]{
				final += j * power_math(len(baseFrom),len(nbr)-i-1)
				break
			}
		}

	}
	return  final
}
func ConvertBase(nbr, baseFrom, baseTo string) string {
	if nbr == "0" && len(baseFrom) > 1 && len(baseFrom) > 1{
		return "0"
	}
	if len(baseFrom) < 2 || len(baseFrom) < 2 {
		return "error"
	}
	final := ""
	dec_num := convert_to_decimal(nbr,baseFrom)
	array_of_index := []int{}
	for dec_num >= 1{
		index := dec_num % len(baseTo)
		array_of_index = append([]int{index},array_of_index...)
		dec_num = dec_num / len(baseTo)
	}

	for i:=0 ; i < len(array_of_index) ; i++ {
		final += string(rune(baseTo[array_of_index[i]]))
	}
	return final

}
func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
	result = ConvertBase("43", "0123456789", "01")
    fmt.Println(result) // 101011
    result = ConvertBase("255", "0123456789", "0123456789ABCDEF")
    fmt.Println(result) // FF
    result = ConvertBase("FF", "0123456789ABCDEF", "0123456789")
    fmt.Println(result) // 255
    result = ConvertBase("77", "01234567", "0123456789")
    fmt.Println(result) // 63
    result = ConvertBase("63", "0123456789", "01234567")
    fmt.Println(result) // 77
    result = ConvertBase("HELLO", "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789")
    fmt.Println(result) // 29234652
    result = ConvertBase("29234652", "0123456789", "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    fmt.Println(result) // HELLO
    result = ConvertBase("0", "0123456789", "01")
    fmt.Println(result) // 0

}
//43
