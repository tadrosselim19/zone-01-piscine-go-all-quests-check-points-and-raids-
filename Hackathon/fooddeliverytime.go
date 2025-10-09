package main

import (
	"fmt"
)

type food struct {
  preptime int
}

func FoodDeliveryTime(order string) int {
	ordermap := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	if item, exist := ordermap[order]; exist {
		return item.preptime
	}
	return 404
}


func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}