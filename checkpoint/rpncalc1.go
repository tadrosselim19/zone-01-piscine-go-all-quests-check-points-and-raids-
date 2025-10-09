package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	array_of_index_to_operator := []int{}
	count_to_index := -1
	array_of_numbers_in_intger := []int{}
	word := ""
	arguments := os.Args[1]
	operators := "+-*/%"
// loop for extraction of operators and numbers in order
	for i:= 0 ; i < len(arguments) ; i++{
		found := false
		for j := 0 ; j < len(operators) ; j++{	
			if arguments[i] == operators[j]{
				count_to_index++
				array_of_index_to_operator = append(array_of_index_to_operator, count_to_index)
				
				array_of_numbers_in_intger= append(array_of_numbers_in_intger, int(operators[j]))
				found = true
			}
		}
		if found == true{
			continue
		}
		if arguments[i] != ' ' && arguments[i] != '\t'{
				word += string(rune(arguments[i]))
		}else{
			if word != ""{
				new_number , err := strconv.Atoi(word)
		        if err != nil{
					fmt.Println("Error")
		       	    return
		        }else{
					count_to_index++
		    	    array_of_numbers_in_intger = append(array_of_numbers_in_intger, new_number)
		        }
			word = ""
			}
		}
	}
fmt.Println(array_of_numbers_in_intger)
fmt.Println(array_of_index_to_operator)

	//printing the result
	if len(array_of_numbers_in_intger) - len(array_of_index_to_operator)!= len(array_of_index_to_operator)+1 {
		fmt.Println("Error")
			return
	}

	for i:=0 ; i < len(array_of_index_to_operator) ; i++{
		if array_of_numbers_in_intger[byte(array_of_index_to_operator[i])] == '+' {
			array_of_numbers_in_intger[array_of_index_to_operator[i]-1] = array_of_numbers_in_intger[array_of_index_to_operator[i]-1-1]  + array_of_numbers_in_intger[array_of_index_to_operator[i]-1]
			array_of_numbers_in_intger = append(array_of_numbers_in_intger[:i], array_of_numbers_in_intger[i+1:]...)
			
		}else if array_of_numbers_in_intger[byte(array_of_index_to_operator[i])] == '-' {
			array_of_numbers_in_intger[array_of_index_to_operator[i]-1] = array_of_numbers_in_intger[array_of_index_to_operator[i]-1-1]  - array_of_numbers_in_intger[array_of_index_to_operator[i]-1]
			array_of_numbers_in_intger = append(array_of_numbers_in_intger[:i], array_of_numbers_in_intger[i+1:]...)
		

			
		}else if array_of_numbers_in_intger[byte(array_of_index_to_operator[i])] == '*' {
			array_of_numbers_in_intger[array_of_index_to_operator[i]-1] = array_of_numbers_in_intger[array_of_index_to_operator[i]-1-1]  * array_of_numbers_in_intger[array_of_index_to_operator[i]-1]
			array_of_numbers_in_intger = append(array_of_numbers_in_intger[:i], array_of_numbers_in_intger[i+1:]...)
			


		}else if array_of_numbers_in_intger[byte(array_of_index_to_operator[i])] == '/'  {
			array_of_numbers_in_intger[array_of_index_to_operator[i]-1] = array_of_numbers_in_intger[array_of_index_to_operator[i]-1-1]  / array_of_numbers_in_intger[array_of_index_to_operator[i]-1]
			array_of_numbers_in_intger = append(array_of_numbers_in_intger[:i], array_of_numbers_in_intger[i+1:]...)
			


		}else if array_of_numbers_in_intger[byte(array_of_index_to_operator[i])] == '%' {
			array_of_numbers_in_intger[array_of_index_to_operator[i]-1] = array_of_numbers_in_intger[array_of_index_to_operator[i]-1-1]  % array_of_numbers_in_intger[array_of_index_to_operator[i]-1]
			array_of_numbers_in_intger = append(array_of_numbers_in_intger[:i], array_of_numbers_in_intger[i+1:]...)
			


		}
		
	}
fmt.Println(array_of_numbers_in_intger)
	fmt.Println(array_of_numbers_in_intger[len(array_of_numbers_in_intger)-1 -1])
fmt.Println(array_of_numbers_in_intger)

	

	

	
}