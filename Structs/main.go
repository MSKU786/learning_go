package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	parth := person{
		firstName: "Parth", 
		lastName: "Amola",
		contactInfo: contactInfo{
			email: "parthamola@gaa.com",
			zipcode: 23422,
		},
	}
	
//parthPointer := &parth; pointer to person is not madatory

// Go can figure out we are tyring to modify person
	parth.updateName("Parthuli");
	parth.print()
}

func (p person ) print () {
	fmt.Printf("%+v", p);
}

func (personToPointer *person) updateName(newFirstName string) {
	*&personToPointer.firstName = newFirstName;
}