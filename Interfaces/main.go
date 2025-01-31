package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb);
	printGreeting(sb);
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting());
}

func (englishBot) getGreeting() string {
	// Customer login specific to english bot
	return "Hi! there"
}

func (spanishBot) getGreeting() string {
	// Customer login specific to spanish bot
	return "Hola!"
}
