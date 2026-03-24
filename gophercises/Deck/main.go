package main

import (
	"fmt"

	"github.com/DistributedSystemAddict/LearnGo/tree/phannam/gophercises/Deck/deck"
)

func main() {
	cards := deck.New(
		deck.WithMultipleDecks(3),
		deck.WithFilter(func(c deck.Card) bool {
			return c.Value == deck.Two || c.Value == deck.Three
		}),
		deck.WithShuffle(),
	)

	fmt.Println("Số lá:", len(cards))
	fmt.Println("5 lá đầu:")
	for _, c := range cards[:5] {
		fmt.Println(" ", c)
	}
}
