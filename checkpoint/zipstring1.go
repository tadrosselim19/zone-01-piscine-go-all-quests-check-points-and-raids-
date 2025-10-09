package main

import (
	"fmt"
)

func convert_to_string(num int)string{
	final := ""
	num_array := []string{}
	if num * -1 > 0{
		num =  num*-1
		final += "-"
	}
	for num >= 1 {
		last := num %10
		num_array = append([]string{string(last + '0')},num_array...)
		num = num / 10 
	}
	for i := 0 ; i < len(num_array) ;i++{
		final += num_array[i]
	}
	return final
}
func ZipString(s string) string {
// sclicing 
	word := ""
	new_s_array := []string{}
	count := 0 
	count_array := []int{}
	// find the unique ch
	for i:= 0 ; i < len(s) ; i++{
		if s[i] != ' ' && s[i] != '\t'{
			word += string(rune(s[i]))
		}else{
			word += string(rune(s[i]))
			if word != ""{
			new_s_array = append(new_s_array, word)
			}
			word = ""
		}
	}
	if word != ""{
		new_s_array = append(new_s_array, word)
		word= ""
	}

  // counting the numbers of each for zipping
    for i:= 0 ; i < len(new_s_array) ; i++{
	    for j := 0 ; j < len(new_s_array[i]) ; j++{
			for j+1 < len(new_s_array[i]) && new_s_array[i][j] == new_s_array[i][j+1]  {
				count++
				j++
			}
			count++
			count_array =append(count_array, count)
			count = 0
		}
	
	}
	
	a:=0
	for i:= 0 ; i < len(new_s_array) ; i++{
	    for j := 0 ; j < len(new_s_array[i]) ; j++{
			word += convert_to_string(count_array[a])
			a++
			word += string(rune(new_s_array[i][j]))
			for j+1 < len(new_s_array[i]) && new_s_array[i][j] == new_s_array[i][j+1]  {
				j++
				continue
			}
			
		}

	}

    return word
}
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
package main

import (
	"fmt"
)

// convert int to []byte (manual, no strconv)
func intToBytes(num int) []byte {
	if num == 0 {
		return []byte{'0'}
	}
	buf := []byte{}
	if num < 0 {
		buf = append(buf, '-')
		num = -num
	}
	digits := []byte{}
	for num > 0 {
		d := byte(num%10) + '0'
		digits = append([]byte{d}, digits...) // prepend
		num /= 10
	}
	buf = append(buf, digits...)
	return buf
}

func ZipString(s string) string {
	var result []byte
	count := 1

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
			continue
		}
		// add number
		result = append(result, intToBytes(count)...)
		// add character
		result = append(result, s[i])
		count = 1
	}
	return string(result)
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

