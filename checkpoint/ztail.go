package main

import "os"

func convert7(num string)(int){
	rune_num := []rune(num)
	final := 0
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[0] == '-'{
			return -1
		}
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return -1 
		}
	}
return final 
}
func main() {
	if len(os.Args) < 4  || os.Args[1] != "-c"{
		return
	}
	// validation of positive number
	pre_number := os.Args[2]
	number := convert7(pre_number)
	if number < 0 {
		return
	}
	files := os.Args[3:]
}