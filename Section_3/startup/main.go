package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Ace of Diamonds"  // redefined, to initialize use :=
	// card := newCard()
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {      // need to give return datatype
	return "Ace of Spades"
}