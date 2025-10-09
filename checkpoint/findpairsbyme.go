package main

import (
	"fmt"
	"os"
	"strconv"
)
func validation_of_array_and_numbers_return(s string)(bool,[]int,string){
    final := []int{}
	pre_final := []string{}
	number_pre := ""
	for i := 0 ;  i < len(s) ; i++{
		if s[0] != '[' || s[len(s)-1] != ']'{
			return false , final , ""
		}
		if s[i] == '['{
			continue
		}
		if s[i] == ']'{
			pre_final= append(pre_final,number_pre)
			number_pre = ""
			break
		}
		if s[i] == ',' {
		   pre_final= append(pre_final,number_pre)
		   number_pre = ""
			continue
		}else if s[i] == ' ' {
			continue
		}else{
			number_pre+= string(rune(s[i]))
		}
		}
	// convert the prenumber to actual number and add them in array
	for i := 0 ; i < len(pre_final) ; i++{
		new_number , err := strconv.Atoi(pre_final[i])
		if err != nil{
			return false,final ,pre_final[i]
		}else{
			final = append(final,new_number)
		}
	}
	return true, final ,""
}
func findpairs(array_of_int []int,num int)[][]int{
	array_of_index :=[][]int{}
	pre_index_array := []int{}
for j := 0 ; j < len(array_of_int) ; j++{
	for i := 0 ; i < len(array_of_int) ; i++{
		if i != j && i > j && array_of_int[i] + array_of_int[j] == num{
			pre_index_array = append(pre_index_array, j,i)
			array_of_index = append(array_of_index, pre_index_array)
			pre_index_array = []int{}
			break
		}else{
			continue
		}
	}
}
	return array_of_index
}
func main(){
	if len(os.Args) != 3{
		fmt.Println("Invalid input.")
		return
	}
	target_number ,err := strconv.Atoi(os.Args[2])
	if err != nil{
		fmt.Println("Invalid target sum.")
		return
	}
	array_of_numbers := os.Args[1]
	// Validaion of characters of arrays given
	for i := 0 ; i <len(array_of_numbers) ;i++{
		if array_of_numbers[0] != '[' || array_of_numbers[len(array_of_numbers)-1] != ']'{
			fmt.Println("Invalid input.")
		return
		}
	}
	flag,new_array,not_number := validation_of_array_and_numbers_return(array_of_numbers)
	if flag == false{
		fmt.Printf("Invalid number: %v\n",not_number)
		return
	}
	if len(findpairs(new_array, target_number)) == 0 {
		fmt.Println("No pairs found.")
		return
	}
	fmt.Printf("Pairs with sum %v: %v\n",target_number,findpairs(new_array,target_number))

}