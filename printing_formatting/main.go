package main

import (
	"fmt"
)

func main() {

	age := 35;
	name := "Ninja"
	// Print doesn't start the new line for new line you have to add \n 
	fmt.Print("Hello .");
	fmt.Print("World !")
	

	// Println always end the line
	fmt.Println("Hello boys!")
	fmt.Println("Let's start with today's go lesson")

	// Below is the example for passing variable in the string
	fmt.Println("My age is", age, "and my name is", name);


	// Printf (formatted Strings) %v = formated specifiers 
	//
	fmt.Printf("my age is %v and my name is %v \n", age, name);
	fmt.Printf("my age is %q and my name is %q \n", age, name);
	fmt.Printf("age is type of %T \n", age);
	fmt.Printf("you scored %f points \n", 223.334);
	fmt.Printf("you scored %0.1f points \n", 223.3233);
	/*
	my age is 35 and my name is Ninja
	my age is '#' and my name is "Ninja"
	age is type of int
	you scored 223.334000 points
	you scored 223.3 points
*/


	//Sprintf => Save formatted string
	var str = fmt.Sprintf("My age is %v and my name is %v", age, name);
	fmt.Println("This is my saved string:", str);
}