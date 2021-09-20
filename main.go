package main

func main() {

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// var cardString = cards.toString()
	// fmt.Println(cardString)
}