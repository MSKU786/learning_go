package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mybill := createBill();

	fmt.Println(mybill);
	
}

func createBill()  bill  {
	reader := bufio.NewReader(os.Stdin);

	// fmt.Print("Crate a bill name: ");
	// name, _ := reader.ReadString('\n');
	// name = strings.TrimSpace(name)

	name,_ := getInput("Create a bill name: ", reader);
	b := newBill(name);
	fmt.Println("Created the bill - ", b.name);
	promptOptions(b);
	return b;
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin);

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t - add tip): ", reader);
	
	switch opt {
		case "a" : 
				name, _ := getInput("Item name.. ", reader);
				price, _ := getInput("Item price.. ", reader);
				p, err := strconv.ParseFloat(price, 64);
				if (err != nil) {
					fmt.Println("This price must be a number");
					promptOptions(b);
				}
				b.addItem(name, p);

				fmt.Println("Item added - ", name, p)
				promptOptions(b);
		case "s" : 
				b.saveBill();
				fmt.Println("you saved the bill in file", b.name);
		case "t" : 
				tip, _ := getInput("Tip ... ", reader);
				t, err := strconv.ParseFloat(tip, 64);
				if (err != nil) {
					fmt.Println("This tip must be a number");
					promptOptions(b);
				}

				b.updateTip(t);
				fmt.Println("Tip Added - ", tip);
				promptOptions(b);
		default: 
				fmt.Println("that was not a valid option")
				promptOptions(b);
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt);

	input, err := r.ReadString('\n');
	input = strings.TrimSpace(input);

	return input, err;
}