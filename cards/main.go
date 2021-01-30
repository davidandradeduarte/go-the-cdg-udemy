package main

import "fmt"

func main() {

	cards := newDeck()

	fmt.Println(cards.saveToFile("my_cards"))

	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	// cards := []string{"Ace of Spades", newCard()}
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	// fmt.Println(card)
	// fmt.Println(cards)
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
