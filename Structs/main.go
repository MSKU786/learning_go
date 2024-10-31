package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	parth := person{firstName: "Parth", lastName: "Amola"}
	fmt.Println(parth);
}