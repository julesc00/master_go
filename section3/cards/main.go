package main

import (
	"fmt"
)

var colors = []string{
	"Red", "Yellow", "Blue", "Orange", "Pink", "Green",
}

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	for i, color := range showColors() {
		fmt.Println(i, color)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func showColors() []string{
	return colors
}
