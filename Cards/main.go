package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Printf(cards.toString());
	hand, remainingCard := deal(cards, 5)

	hand.print()
	remainingCard.print()
}
