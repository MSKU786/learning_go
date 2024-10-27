package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//create a new data type 'deck'
// which is basically a slice of strings

type deck []string

func newDeck () deck {
		cards := deck{};

		cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
		cardNums:= []string {"Ace", "Two", "Three", "Four"}

		for _, suit := range cardSuits {
			for _, value := range cardNums {
				cards = append(cards, value + " of "+ suit);
			}
		}

		return cards;
}

func (d deck) print () {
	for i, card := range(d ){
		fmt.Println(i, card);
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) toString () string  {
		deckString := strings.Join([]string(d), ",");
		return deckString
}

func ( d deck) saveToFile (fileName string) error {
		return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile (filename string) deck {
		bs, err := os.ReadFile(filename)

		if (err != nil) {
			fmt.Println("Error:", err)
			os.Exit(1);
		}

		s := strings.Split(string(bs),",")

		return deck(s);
}


func (d deck) shuffel () {
		for i := range (d) {
				randInd := rand.Intn(len(d)-1);
				d[i], d[randInd] = d[randInd], d[i]
		}
}