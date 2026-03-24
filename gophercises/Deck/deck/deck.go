package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Option func([]Card) []Card

var defaultSuits = []Suit{Spades, Diamonds, Clubs, Hearts}

func New(opts ...Option) []Card {
	cards := make([]Card, 0, 52)
	for _, suit := range defaultSuits {
		for value := Ace; value <= King; value++ {
			cards = append(cards, Card{Suit: suit, Value: value})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func WithSort(less func(i, j int) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less)
		return cards
	}
}

func DefaultLess(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit < cards[j].Suit
		}
		return cards[i].Value < cards[j].Value
	}
}

func WithShuffle() Option {
	return func(c []Card) []Card {
		r := rand.New(rand.NewSource((time.Now().UnixNano())))
		r.Shuffle(len(c), func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
		return c
	}
}

func WithJokers(n int) Option {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

func WithFilter(filter func(Card) bool) Option {
	return func(cards []Card) []Card {
		result := make([]Card, 0, len(cards))
		for i := 0; i < len(cards); i++ {
			if !filter(cards[i]) {
				result = append(result, cards[i])
			}
		}
		return result
	}
}

func WithMultipleDecks(n int) Option {
	return func(cards []Card) []Card {
		result := make([]Card, 0, len(cards)*n)
		for i := 0; i < n; i++ {
			for _, suit := range defaultSuits {
				for value := Ace; value <= King; value++ {
					result = append(result, Card{Suit: suit, Value: value})
				}
			}
		}
		return result
	}
}
