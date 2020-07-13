package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
