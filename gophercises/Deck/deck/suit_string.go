package deck

import "strconv"

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Hearts:
		return "Hearts"
	case Joker:
		return "Joker"
	default:
		return "Suit(" + strconv.Itoa(int(s)) + ")"
	}
}

func (v Value) String() string {
	switch v {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Value(" + strconv.Itoa(int(v)) + ")"
	}
}
