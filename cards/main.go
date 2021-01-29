package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for _, card := range cards {
		fmt.Println(card)
	}

	// fmt.Println(card)
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
