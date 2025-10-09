package main

import "fmt"

func Sqrt(nb int) int {
	if nb <= 0{
		return 0
	}
	for i:= 0 ; i < nb ; i++{
		if i * i == nb{
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(64))
	fmt.Println(Sqrt(3))
}

func aSqrt(nb int) int {
	result :=[]int{}
	for i := nb ; i >= 2 ; i = i / 2 {
		result = append(result, i)
	}
	return len(result)


}
func bSqrt(nb int) int {
	for i := 1 ; i <= nb ; i++{
		if i * i == nb{
			return i
		}
	}
	return 0

}
