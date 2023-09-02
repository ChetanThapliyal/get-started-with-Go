package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if value, exists := cardValues[card]; exists {
		return value
	}
	return 0 // Default value if card is not found
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P" // Split pair of aces
	case sum == 21 && (dealerCard != "ace" && dealerCard != "king" && dealerCard != "queen" && dealerCard != "jack"):
		return "W" // Automatically win with Blackjack if dealer doesn't have high cards
	case sum >= 17 && sum <= 20:
		return "S" // Stand when cards sum up to 17-20
	case sum >= 12 && sum <= 16 && (ParseCard(dealerCard) >= 7):
		return "H" // Hit when cards sum up to 12-16 and dealer has 7 or higher
	case sum <= 11:
		return "H" // Hit when cards sum up to 11 or lower
	default:
		return "S" // Stand for any other scenario
	}
}
