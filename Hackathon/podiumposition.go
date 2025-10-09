package main

import (
	"fmt"
)
func PodiumPosition(podium [][]string) [][]string {
	for i :=0 ; i < len(podium) ; i++{
		for j :=0 ; j < len(podium) ; j++{
		if podium[i][0] < podium[j][0]{
			podium[i] , podium[j] = podium[j] , podium[i]
		}
		}
	}
	return podium
}
func aPodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	var result [][]string
	if n == 1 {
		result = [][]string{{podium[0][0]}}
	} else if n == 2 {
		result = [][]string{podium[1], podium[0]}
	} else if n == 3 {
		result = [][]string{podium[2], podium[1], podium[0]}
	} else if n == 4 {
		result = [][]string{podium[3], podium[2], podium[1], podium[0]}
	}
	return result
}
func main() {

    p4 := []string{"4th Place"}
    p3 := []string{"3rd Place"}
    p2 := []string{"2nd Place"}
    p1 := []string{"1st Place"}

    position := [][]string{p4, p3, p2, p1}
    fmt.Println(PodiumPosition(position))
}