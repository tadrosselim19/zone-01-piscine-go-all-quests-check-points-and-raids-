package main

import "fmt"


func IterativePower(nb int, power int) int {
	if power < 0{
		return 0
	}
	result := 1
	for i := 1; i <= power; i++ {
		result *= nb
	}
	return result
}


func main() {
	fmt.Println(IterativePower(-6,3))

}