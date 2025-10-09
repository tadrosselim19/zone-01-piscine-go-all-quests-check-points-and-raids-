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
	for i := 0; i < len(arguments); i++ {
		found := false
		for j := 0; j < len(operators); j++ {
			if arguments[i] == operators[j] {
				count_to_index++
				array_of_index_to_operator = append(array_of_index_to_operator, int(operators[j]))
				array_of_numbers_in_intger = append(array_of_numbers_in_intger, int(operators[j]))
				found = true
			}
		}
		if found {
			continue
		}
		if arguments[i] != ' ' && arguments[i] != '\t' {
			word += string(rune(arguments[i]))
		} else {
			if word != "" {
				new_number, err := strconv.Atoi(word)
				if err != nil {
					fmt.Println("Error")
					return
				} else {
					count_to_index++
					array_of_numbers_in_intger = append(array_of_numbers_in_intger, new_number)
				}
				word = ""
			}
		}
	}
	if word != "" {
		new_number, err := strconv.Atoi(word)
		if err != nil {
			fmt.Println("Error")
			return
		} else {
			count_to_index++
			array_of_numbers_in_intger = append(array_of_numbers_in_intger, new_number)
		}
	}

	// evaluation
	stack := []int{}
	for i := 0; i < len(array_of_numbers_in_intger); i++ {
		val := array_of_numbers_in_intger[i]

		if val == '+' || val == '-' || val == '*' || val == '/' || val == '%' {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}
			a := stack[len(stack)-2] // second last
			b := stack[len(stack)-1] // last
			stack = stack[:len(stack)-2]

			if val == '+' {
				stack = append(stack, a+b)
			} else if val == '-' {
				stack = append(stack, a-b) // ✅ correct order
			} else if val == '*' {
				stack = append(stack, a*b)
			} else if val == '/' {
				if b == 0 {
					fmt.Println("Error")
					return
				}
				stack = append(stack, a/b) // ✅ correct order
			} else if val == '%' {
				if b == 0 {
					fmt.Println("Error")
					return
				}
				stack = append(stack, a%b) // ✅ correct order
			}
		} else {
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}
