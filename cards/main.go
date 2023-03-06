package main

import "fmt"

func main() {

	cards := newDeck()
	cards.shuffle()
	fmt.Println(len(cards))
	// cards.print()
	// cards := newDeckFromFile("My_Hand_Cards")
	// cards.print()
	// fmt.Println(cards)
	// cards := newDeck()

	// hand, _ := deal(cards, 5)
	// hand.saveToFile("My_Hand_Cards")
	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("save deck to file")
	// fmt.Println(cards.toString())
}
