package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <2 {
		return
	}
	arguments := os.Args[1:]
	

	for i := 0 ; i < len(arguments) ; i++{
		array_of_index := []int{}
		flag_of_parentheses := 0
	    flag_of_square := 0 
	    flag_of_curly := 0 
		generalized_flag := true
		if arguments[i] == ""{
			z01.PrintRune('O')
			z01.PrintRune('k')
			z01.PrintRune('\n')
			continue
		}
		// count number of each flage
		for j := 0 ; j < len(arguments[i]) ; j++{
			if j != len(arguments[i]) && ((arguments[i][j] == '(' && arguments[i][j+1] == ')') || (arguments[i][j] == '[' && arguments[i][j+1] == ']') || (arguments[i][j] == '{' && arguments[i][j+1] == '}')){
				j++
				continue
			}
			if arguments[i][j] == '(' {
				flag_of_parentheses++
				array_of_index = append(array_of_index, 1)
			}else if arguments[i][j] == '{'{
				flag_of_curly++
				array_of_index = append(array_of_index, 3)
			}else if arguments[i][j] == '['{
				flag_of_square++
				array_of_index = append(array_of_index, 2)
			}else if arguments[i][j] == ')' {
				flag_of_parentheses--
				array_of_index = append(array_of_index, 4)
			}else if arguments[i][j] == '}'{
				flag_of_curly--
				array_of_index = append(array_of_index, 6)
			}else if arguments[i][j] == ']'{
				flag_of_square--
				array_of_index = append(array_of_index, 5)
			}
		}
		//fmt.Println(array_of_index)
		// checking the arrangment of barackets
		for j := 0 ; j < len(array_of_index)/2 ; j++{
			for k := len(array_of_index)-1 ; k >= len(array_of_index)/2 ; k--{
				if array_of_index[j] + 3 == array_of_index[k]{
					generalized_flag = true
				}else{
					generalized_flag = false
				}
			}
		}

		// printing the result of argument 
		if generalized_flag == false || flag_of_parentheses != 0 || flag_of_square != 0 || flag_of_curly != 0{
			z01.PrintRune('E')
			z01.PrintRune('r')
			z01.PrintRune('r')
			z01.PrintRune('o')
			z01.PrintRune('r')
			z01.PrintRune('\n')
		}else{
			z01.PrintRune('O')
			z01.PrintRune('k')
			z01.PrintRune('\n')
		}
	}
}