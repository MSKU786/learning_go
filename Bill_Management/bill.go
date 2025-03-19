package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//Make new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}


// Format the bill (reciever function)
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	total := 0.0

	// List Items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v", k+":", v)
		total += v
	}

	// Total
	fs += fmt.Sprintf("-25v ...$%0.2f", "Total:", total);
	return fs;
}

// UPdate tip
func (b *bill) updateTip(tip float64) {
	// Go will automatically resolved this we don't have mention specifically (*b).tip = tip. 
	// Struct and pionters are automaticlaly derefernce
	b.tip = tip;
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price;
}