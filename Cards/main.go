package main

func main() {
	cards := deck{"Ace fo Diamond", newCard()}
	cards = append(cards, "Five of Spade")
	cards.print()
}

func newCard() string {
	return "Five of Diamond"
}