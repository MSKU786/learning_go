package main

import "fmt"

//create a new data type 'deck'
// which is basically a slice of strings

type deck []string

func newDeck () deck {
		cards := deck{};

		cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
		cardNums:= []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Joker", "Queen", "King"}

		for _, suit := range cardSuits {
			for _, value := range cardNums {
				cards = append(cards, value + " f "+ suit);
			}
		}

		return cards;
}

func (d deck) print () {
	for i, card := range(d ){
		fmt.Println(i, card);
	}
}