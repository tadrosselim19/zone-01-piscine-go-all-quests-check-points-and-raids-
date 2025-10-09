package main

import "fmt"

func defineRectangle(ptr *point, n int) *rectangle {

}

func calArea(ptr *rectangle) int {

}

func main() {
	vPoint := []point{}
	rectangle := &rectangle{}
	n := 7

	for i := 0; i < n; i++ {
		val := point{
			x: i%2 + 1,
			y: i + 2,
		}
		vPoint = append(vPoint, val)
	}
	rectangle = defineRectangle(vPoint, n)
	fmt.Println("area of the rectangle:", calArea(rectangle))
}