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
	case "ten", "jack", "queen", "king":
		return 10
	}
	/*
		| card  | value | card    | value |
		| :---: | :---: | :-----: | :---: |
		|  ace  |  11   | eight   |   8   |
		|  two  |   2   | nine    |   9   |
		| three |   3   |  ten    |  10   |
		| four  |   4   | jack    |  10   |
		| five  |   5   | queen   |  10   |
		|  six  |   6   | king    |  10   |
		| seven |   7   | *other* |   0   |
	*/
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerInt := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum >= 20:
		if dealerInt >= 10 {
			return "S"
		} else {
			return "W"
		}

	case sum >= 17 && sum <= 20:
		return "S"

	case sum >= 12 && sum <= 16:
		if dealerInt >= 7 {
			return "H"
		} else {
			return "S"
		}
	}

	return "H"

	/*	Colinhas
		- If you have a pair of aces you must always split them.

		- If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.

		- If your cards sum up to a value within the range [17, 20] you should always stand.

		- If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.

		- If your cards sum up to 11 or lower you should always hit.

		- Stand (S)
		- Hit (H)
		- Split (P)
		- Automatically win (W)
	*/
}
