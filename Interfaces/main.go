package main

type englishBot struct {}
type spanishBot struct {}

func main() {

}

func (eb englishBot) getGreeting () string {
		// Customer login specific to english bot
		return "Hi! there"
}


func (eb spanishBot) getGreeting () string {
	// Customer login specific to spanish bot
	return "Hola!"
}

