package main

import "fmt"

// create a new type of 'deck'
// which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// d is the reference of cards in main.go file
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
