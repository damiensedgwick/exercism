package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	pairOfAces := card1 == "ace" && card2 == "ace"
	blackJack := value == 21
	dealHasHighCard := dealerValue >= 10
	dealerHasASevenOrHigher := dealerValue >= 7

	switch {
	case pairOfAces:
		return "P"
	case blackJack && dealHasHighCard:
		return "S"
	case blackJack && !dealHasHighCard:
		return "W"
	case value >= 17 && value <= 20:
		return "S"
	case value >= 12 && value <= 16 && dealerHasASevenOrHigher:
		return "H"
	case value >= 12 && value <= 16 && !dealerHasASevenOrHigher:
		return "S"
	case value <= 11:
		return "H"
	default:
		return "S"
	}
}
