package main

import "fmt"


func DealAPackOfCards(deck []int) {
	everyplayer := len(deck) / 4
	for j := 0; j < 4; j++ {
		start := j * everyplayer
		end := start + everyplayer
		playercard := deck[start:end]

		fmt.Printf("Player %d: ", j+1)
		for i, card := range playercard {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(card)
		}
		fmt.Println()
	}
}

func aDealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: %d, %d, %d\n", i+1, deck[i*3], deck[i*3+1], deck[i*3+2])
	}
}


func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}