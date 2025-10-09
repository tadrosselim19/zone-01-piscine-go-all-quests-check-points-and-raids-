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

		array_of_oppened_bracket:=[]rune{}
		if arguments[i] == ""{
			z01.PrintRune('O')
			z01.PrintRune('k')
			z01.PrintRune('\n')
			continue
		}
		validation := true

	for j := 0 ; j < len(arguments[i]) ; j++{
		if arguments[i][j]== '(' || arguments[i][j]== '[' || arguments[i][j]== '{' {
			array_of_oppened_bracket = append(array_of_oppened_bracket, rune(arguments[i][j]))

		}else if arguments[i][j] == ')'{
			if len(array_of_oppened_bracket) == 0 || array_of_oppened_bracket[len(array_of_oppened_bracket)-1] != '('{
				validation = false
				break
			}else{
				array_of_oppened_bracket = array_of_oppened_bracket[0:len(array_of_oppened_bracket)-1]
			}
		}else if arguments[i][j] == ']'{
			if len(array_of_oppened_bracket) == 0 || array_of_oppened_bracket[len(array_of_oppened_bracket)-1] != '['{
				validation = false
				break
			}else{
				array_of_oppened_bracket = array_of_oppened_bracket[0:len(array_of_oppened_bracket)-1]
			}
		}else if arguments[i][j] == '}'{
			if len(array_of_oppened_bracket) == 0 || array_of_oppened_bracket[len(array_of_oppened_bracket)-1] != '{'{
				validation = false
				break
			}else{
				array_of_oppened_bracket = array_of_oppened_bracket[0:len(array_of_oppened_bracket)-1]
			}
		}

	}
	// printing 
	if validation == false || len(array_of_oppened_bracket) !=0{
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