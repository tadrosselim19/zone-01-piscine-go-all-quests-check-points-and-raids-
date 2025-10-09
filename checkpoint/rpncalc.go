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
	array_of_numbers_in_string := []string{}
	word := ""
	opperators_found := []byte{}
	arguments := os.Args[1]
	operators := "+-*/%"
// loop for extraction of operators and numbers in order
	for i:= 0 ; i < len(arguments) ; i++{
		found := false
		for j := 0 ; j < len(operators) ; j++{	
			if arguments[i] == operators[j]{
				opperators_found= append(opperators_found, operators[j])
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
			array_of_numbers_in_string = append(array_of_numbers_in_string, word)
			word = ""
			}
		}
	}
// fmt.Println(opperators_found)

// // loop for rearange the values of opperator
//     for i := 0 ; i < len(opperators_found) ; i++{
// 		for j := 0 ; j < len(opperators_found) ; j++{
// 			if opperators_found[i] < opperators_found[j]{
// 				opperators_found[i] , opperators_found[j] = opperators_found[j] , opperators_found[i]
// 			}
// 		}
// 	}
// fmt.Println(opperators_found)

// convertion of numbers from string to integer
    array_of_numbers_in_intger := []int{}
	
	for i:= 0 ; i < len(array_of_numbers_in_string) ; i++{
		new_number , err := strconv.Atoi(array_of_numbers_in_string[i])
		if err != nil{
			fmt.Println("Error")
			return
		}else{
			array_of_numbers_in_intger = append(array_of_numbers_in_intger, new_number)
		}
	}

	// printing the result
	if len(opperators_found) + 1 != len(array_of_numbers_in_intger){
		fmt.Println("Error")
			return
	}

	for i := 0 ; i < len(opperators_found);i++{
		if opperators_found[i] == '+'{
			array_of_numbers_in_intger[i+1] = array_of_numbers_in_intger[i] + array_of_numbers_in_intger[i+1]
		}else if opperators_found[i] == '-'{
			array_of_numbers_in_intger[i+1] = array_of_numbers_in_intger[i] - array_of_numbers_in_intger[i+1]
		}else if opperators_found[i] == '*'{
			array_of_numbers_in_intger[i+1] = array_of_numbers_in_intger[i] * array_of_numbers_in_intger[i+1]
		}else if opperators_found[i] == '/' && array_of_numbers_in_intger[i+1] != 0{
			array_of_numbers_in_intger[i+1] = array_of_numbers_in_intger[i] / array_of_numbers_in_intger[i+1]
		}else if opperators_found[i] == '%' && array_of_numbers_in_intger[i+1] != 0{
			array_of_numbers_in_intger[i+1] = array_of_numbers_in_intger[i] % array_of_numbers_in_intger[i+1]
		}else{
			fmt.Println("Error")
			return
		}

	}
	
	fmt.Println(array_of_numbers_in_intger[len(array_of_numbers_in_intger)-1])
	

	
}