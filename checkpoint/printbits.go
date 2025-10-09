package main

import (
	
	"os"
)
func convert5(num string)(int,bool){
	rune_num := []rune(num)
	final := 0
	flag := 1
	for i := 0 ;i < len(rune_num); i++{
		if rune_num[0] == '-'{
			flag *= -1
			continue
		}
		if rune_num[i] <= '9' && rune_num[i] >= '0'{
			final = final *10 + int(rune_num[i])-'0'
		}else{
			return -1 ,false
		}
	}
return final * flag ,true
}

func serch_for_highest_power(a int)int{
	b:= a
	for b > 0{
		if b == 1{
			return a
		}
		if b %2 == 0{
			b = b /2
		}else{
			a--
			b = a
		}
	   }
	return 0
}
func main() {
	if len(os.Args) != 2{
		return
	}
	number , err := convert5(os.Args[1])
	if err == false{
		os.Stdout.WriteString("invaled input")
		return
	}
	changable_number := number
	b := serch_for_highest_power(changable_number)
	sorting_bit := []int{}
	compartion_bit := []int{}
	for changable_number > 0{
	    a := serch_for_highest_power(changable_number)
		sorting_bit = append(sorting_bit, a)
		changable_number = changable_number - a
	}
	for i := b ; i >= 1 ; i/=2{
		compartion_bit = append(compartion_bit, i)
	}

	// fmt.Println(sorting_bit)
	// fmt.Println(compartion_bit)

	// adjustment on comparisonbit to fit represented bit
	for i := 0 ; i < len(sorting_bit) ; i++{
		for j := 0 ; j < len(compartion_bit) ; j++{
			if compartion_bit[j] == sorting_bit[i]{
				compartion_bit[j]= 3
				break
			}
		}

	}
	for i:= 0 ; i < len(compartion_bit) ; i++{
		if compartion_bit[i] != 3{
			compartion_bit[i] =0 
		}else{
			compartion_bit[i]= 1
		}
	}
	// puting numbers in its bit
	represented_4_byte := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	i :=len(compartion_bit)-1
		for j :=len(represented_4_byte)-1 ; i >= 0  ; j--{
			
			represented_4_byte[j] = compartion_bit[i]
			i--
		}
	

	// print the final result 
	count := 0
	for i:=0 ; i < len(represented_4_byte) ; i++{
		if count != 8 {
		os.Stdout.WriteString(string(represented_4_byte[i]+'0'))
		count++
		}else{
			os.Stdout.WriteString(" ")
			os.Stdout.WriteString(string(represented_4_byte[i]+'0'))
			count = 1
		}
	}
}

//notes
// Decimal to Binary Example
// Convert decimal 25 to binary:
// 1- Find largest power of 2 ≤ 25 → 16 (2⁴)
// 2- Subtract: 25 - 16 = 9
// 3- Largest power of 2 ≤ 9 → 8 (2³)
// 4- Subtract: 9 - 8 = 1
// 5- Largest power of 2 ≤ 1 → 1 (2⁰)