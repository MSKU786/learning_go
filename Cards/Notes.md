# Defining variables

var card string = "adsf adf" //long form
card := "adfa " // short form

both syntax carry equal weight , in first syntar we defining a new variable which will contain string value

# defining function and return types

func newCard() string {
return "Five of Diamonds"
}

newCard() is the function defination and declartion next paramter string is the return value

# Slice Range Syntax:

fruits := []string {"apple", "banana", "grape", "oragne"}
fruits[startIndexIncluding: upToNotIncluding]

# Two ways to define methods

## Method version:

func (d deck) deal(handSize int) (deck, deck) {
return d[:handSize], d[handSize:]
}

## Function version:

func deal(d deck, handSize int) (deck, deck) {
return d[:handSize], d[handSize:]
}

In Go, there are two ways to define methods on types:

Method with receiver: The syntax func (d deck) is for defining methods, where d is the receiver. Methods are tied to specific types.

Regular function: The second version is a regular function that takes arguments, not a method tied to a specific type.
