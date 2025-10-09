package main

import (
	"fmt"
	"os"
	"strconv"
)
func to_roman(s []string,n []int , num int)(string,string){
	final1 := ""
	final2 := ""
	for num != 0{
		for i := len(n)-1 ; i >=0 ; i--{
			if num - n[i] > 0 && num - n[i]!=0{
				final1 += s[i]
				if len(s[i]) >1{
					final2 += "(" + string(rune(s[i][1])) + "-" + string(rune(s[i][0])) + ")" + "+"
				}else{
					final2 += s[i] + "+"
				}
				
				num = num - n[i]
				break
			}else if num - n[i] == 0 {
				final1 += s[i]
				if len(s[i]) >1{
					final2 += "(" + string(rune(s[i][1])) + "-" + string(rune(s[i][0])) + ")"
				}else{
					final2 += s[i]
				}
				num = num - n[i]
				break
			}else{
				continue
			}
		}
		
	}
	return final1 ,final2
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	number , err := strconv.Atoi(os.Args[1])
	if err != nil || number>= 4000 || number <= 0{
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	roman_list := []string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}
	roman_equ := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
	roman_num ,eqq := to_roman(roman_list,roman_equ,number)
	fmt.Println(eqq)
	fmt.Println(roman_num)
}