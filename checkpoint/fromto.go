package main

import (
	"fmt"
)

func FromTo(from int, to int)string{

	final := ""

	if from > 99 || to > 99 || from < 0 || to < 0 {
		return "Invalid\n"
	}

	if from <= to{
		for i := from ; i <= to ; i++{
			if i != to {
				final += convert(i) + ", "
			}else{
				final += convert(i)
			}
	    }
	}else {
		for i := from ; i >= to ; i--{
			if i != to {
				final += convert(i) + ", "
			}else{
				final += convert(i) 
			}
		}
	}
return final + "\n"
}
func convert(num int)string{
	if num <= 9 {
		return "0" + string(rune(num+'0')) 
	}
	num_string := ""
	b := num
	to_array_int := []int{}
	for b > 0 {
		to_array_int = append(to_array_int, b%10) 
		b= b /10
    }
	for i:= len(to_array_int)-1 ; i >= 0 ; i--{
		num_string += string(rune(to_array_int[i]+'0'))
	}
	return num_string
}


func main(){
	fmt.Print(FromTo(98,88))
}