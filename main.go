package main

func main() {

	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards := []string{newCard(), "Ace Of Spades"}
	// cards := deck{newCard(), "Ace Of Spades"}

	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	// for i, card := range cards {
	// }
	// for i, card := range cards {
	// 	fmt.Println(card)
	// 	i = i + 1
	// }
	// cards.print()

	//byte conversion
	// tobe := "The String"
	// fmt.Println([]byte(tobe))
	// }

	// func newCard() string {
	// 	return "Five Of Diamonds"
}
