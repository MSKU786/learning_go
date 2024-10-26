package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Printf(cards.toString());
	cards.saveToFile("cards.txt");
	hand, remainingCard := deal(cards, 5)

	hand.print()
	remainingCard.print()
}
