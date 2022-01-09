package main

func main() {
	// var card string = "Ace of Spades"
	//cards := newDeck() //:= Use when assigning new variable, first time declaration
	//cards.saveToFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
}
