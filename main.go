package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remCards := deal(cards, 5)
	// hand.print()
	// remCards.print()
	cards := newDeck()
	cards.shuffle()
	fmt.Println(cards)
	// cards.saveToFile("poker_cards")
	// fmt.Println(readFile("poker_cards"))
}
