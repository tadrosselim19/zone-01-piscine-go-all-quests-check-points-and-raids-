package main

import (
	"fmt"
	
)
func ShoppingListSort(slice []string) []string {
	for i := 0 ; i < len(slice) ; i++{
		for j := 0 ; j < len(slice) ; j++{
			if len(slice[i]) < len(slice[j]){
				slice[i] ,slice[j] =  slice[j] , slice[i]
			}
		}
	}
	return slice
}
func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}