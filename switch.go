package main

import "fmt"

func ParseCard(name string) int {
	var value = 0
	switch name {
	case "one":
		value = 1
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	case "ace":
		value = 11
	}
	return value
}

func FormatCardName(card_name, card_type string, value int) string {
	return fmt.Sprintf("%s is %s and is worth %d", card_name, card_type, value)
}

/*
If you have a pair of aces you must always split them.
If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win.

	If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.

If your cards sum up to a value within the range [17, 20] you should always stand.
If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
If your cards sum up to 11 or lower you should always hit.
*/
func FirstTurn(card1, card2, dealer_card string) string {

	var card1_value = ParseCard(card1)
	fmt.Println(FormatCardName("card1", card1, card1_value))
	var card2_value = ParseCard(card2)
	fmt.Println(FormatCardName("card2", card2, card2_value))
	var dealer_card_value = ParseCard(dealer_card)
	fmt.Println(FormatCardName("dealer card", dealer_card, dealer_card_value))
	var sum = card1_value + card2_value
	var action = "NULL"
	// Two aces -> Split
	if card1_value == 11 && card2_value == 11 {
		action = "P"
		return action
	}
	switch sum {
	case 12, 13, 14, 15, 16:
		if dealer_card_value >= 7 {
			action = "H"
		} else {
			action = "S"
		}
	case 17, 18, 19, 20:
		action = "S"
	case 21:
		if dealer_card_value >= 10 {
			action = "S"
		} else {
			action = "W"
		}
	default:
		action = "H"
	}
	return action
}

func SwitchTest() {
	var card1_name = "ace"
	fmt.Printf("my %s is worth %d\n", card1_name, ParseCard(card1_name))
	var card2_name = "ace"
	var dealer_card_name = "jack"
	fmt.Println("player should", FirstTurn(card1_name, card2_name, dealer_card_name))
	card2_name = "king"
	dealer_card_name = "ace"
	fmt.Println("player should", FirstTurn(card1_name, card2_name, dealer_card_name))
	card1_name = "five"
	card2_name = "queen"
	fmt.Println("player should", FirstTurn(card1_name, card2_name, dealer_card_name))
}
