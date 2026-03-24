package deck

import "fmt"

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

type Value int

const (
	Ace   Value = iota + 1 // 1
	Two                    // 2
	Three                  // 3
	Four                   // 4
	Five                   // 5
	Six                    // 6
	Seven                  // 7
	Eight                  // 8
	Nine                   // 9
	Ten                    // 10
	Jack                   // J
	Queen                  // Q
	King
)

type Card struct {
	Suit  Suit
	Value Value
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

// Trong Go, nếu một kiểu có method tên String(),
// thì bất cứ khi nào dùng fmt.Println()
// hay fmt.Sprintf("%s", ...) với kiểu đó,
// Go sẽ tự động gọi method String() này. Đây là interface fmt.Stringer.
