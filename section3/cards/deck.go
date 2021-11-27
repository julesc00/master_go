package main

import "fmt"

// Create a new type of "deck", which is a slice of strings.
// Custom type declarations
type deck []string

// Receiver on a function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}