package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	// appending elemts to slice
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of diamonds"
}
