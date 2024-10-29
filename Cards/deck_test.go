package main

import "testing"

func TestNewDec(t *testing.T) {
		d := newDeck();

		if (len(d) != 16) {
				t.Errorf("Excpected deck length is 16 but we got %v", len(d));
		}
}