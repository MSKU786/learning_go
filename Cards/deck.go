package main

import "fmt"

//create a new data type 'deck'
// which is basically a slice of strings

type deck []string

func (d deck) print () {
	for i, card := range(d ){
		fmt.Println(i, card);
	}
}