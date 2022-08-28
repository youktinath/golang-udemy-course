package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 10)

	// fmt.Println("The HAND:")
	// hand.print()
	// fmt.Println("The OTHERS:")
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("the_deck.txt")
	// cardsNew := newDeckFromFile("the_deck.txt")
	// cardsNew.print()
	cards.shuffle()
	cards.print()
}