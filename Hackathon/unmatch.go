package main

import (
	"fmt"
)
func Unmatch(a []int) int {
	for i := 0 ; i < len(a) ; i ++{
		count := 0
            for j := 0 ; j < len(a) ; j ++{
		        if a[i] == a[j] {
				count ++
			   }
	        }
		if count%2 != 0{
			return a[i] 
		}

    }
 return -1
}
func aUnmatch(a []int) int {
	counts := make(map[int]int)

	for _, num := range a {
		counts[num]++
	}

	for _, num := range a {
		if counts[num]%2 != 0 {
			return num
		}
	}

	return -1
}
func UnmatchAll(a []int) []int {
	counts := make(map[int]int)
	for _, v := range a {
		counts[v]++
	}

	var result []int
	for num, cnt := range counts {
		if cnt%2 != 0 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4 , 5}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}