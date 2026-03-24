package deck

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Errorf("Mong đợi 52 lá, nhưng nhận được %d lá", len(cards))
	}
}

func TestNewWithJokers(t *testing.T) {
	cards := New(WithJokers(2))
	if len(cards) != 54 {
		t.Errorf("Mong đợi 54 lá, nhưng nhận được %d lá", len(cards))
	}
}

func TestNewWithFilter(t *testing.T) {
	filter := func(c Card) bool {
		return c.Value == Two || c.Value == Three
	}
	cards := New(WithFilter(filter))
	// 52 - 4 lá 2 - 4 lá 3 = 44 lá
	if len(cards) != 44 {
		t.Errorf("Mong đợi 44 lá, nhưng nhận được %d lá", len(cards))
	}
}

func TestNewWithMultipleDecks(t *testing.T) {
	cards := New(WithMultipleDecks(3))
	if len(cards) != 156 {
		t.Errorf("Mong đợi 156 lá, nhưng nhận được %d lá", len(cards))
	}
}

func TestNewWithSort(t *testing.T) {
	cards := New(WithShuffle(), WithSort(DefaultLess(New())))
	if cards[0] != (Card{Suit: Spades, Value: Ace}) {
		t.Errorf("Lá đầu tiên phải là Ace of Spades, nhưng nhận được %s", cards[0])
	}
}

func ExampleNew() {
	cards := New()
	for _, c := range cards[:5] {
		fmt.Println(c)
	}
}
