package main

import (
	"fmt"
)

type Item struct{
	name string
	amount float32
}

type Bill struct {
	bill_name string
	items Item
}

func main() {
	var bill_name string;
	fmt.Printf("Create a bill name:");
	fmt.Scanln(&bill_name)
	fmt.Printf("Create the bill - %v", bill_name );



}

func addItem() {
	while(1) {
		var option string;
		var option_picked string;
		fmt.Printf("Choose option (a - add item, s - save bill, t - add tip) : ")
		fmt.Scanln(&option_picked);
	
		switch(option_picked) {
			case 'a' : 
				option = 'add item'
		}
	}
}