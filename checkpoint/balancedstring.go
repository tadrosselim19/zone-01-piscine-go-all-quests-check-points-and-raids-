package main

import "fmt"

func balanced_string(s string) int {
	count_c := 0
	count_d := 0
	count_both := 0 
	for i := 0 ; i < len(s) ; i++{
		if s[i] != 'D' && s[i] != 'C'{
			return 0
		}

		if s[i] == 'C'{
			count_c++

	    }
		if s[i] == 'D'{
			count_d++

	    }
		if count_c == count_d{
			count_both++
			count_c = 0 
			count_d = 0
		}
	}
	return count_both
}
func main() {
	fmt.Println(balanced_string("CDCCDDCDCD"))
	fmt.Println(balanced_string("CDDDDCCCDC"))
	fmt.Println(balanced_string("DDDDCCCC"))
	fmt.Println(balanced_string("CDCCCDDCDD"))

}