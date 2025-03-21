package main

import (
	"bufio"
	"fmt"
	"os"
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

	return b;
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin);

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t - add tip): ", reader);
	fmt.Println(opt);
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt);

	input, err := r.ReadString('\n');
	input = strings.TrimSpace(input);

	return input, err;
}