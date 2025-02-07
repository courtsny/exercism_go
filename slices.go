package main

import (
	"fmt"
	"slices"
)

// No const needed since arrays are immutable
var favorite_cards = [3]int{2, 6, 9}

func FavoriteCards() []int {
	return favorite_cards[:]
}

// If found, return index card is at
// Else, return -1
func GetItem(cards []int, card_to_find int) int {
	for index, element := range cards {
		if card_to_find == element {
			return index
		}
	}
	return -1
}

func SetItem(cards []int, index, new_card int) []int {
	if index >= len(cards) || index < 0 {
		cards = append(cards, new_card)
	} else {
		cards[index] = new_card
	}
	return cards
}

func PrependItems(cards []int, new_cards ...int) []int {
	cards = slices.Concat(new_cards, cards)
	return cards
}

func RemoveItem(cards []int, card_to_remove int) []int {
	var index = GetItem(cards, card_to_remove)
	if index > -1 {
		cards = slices.Concat(cards[:index], cards[index+1:])
	} else {
		fmt.Println("card", card_to_remove, "DNE in the deck")
	}
	return cards
}

func SlicesTest() {
	// Favorite
	fmt.Println("favorite cards are:", favorite_cards)
	// Search
	card_to_find := 2
	cards := []int{3, 7, 8, 1, 4}
	fmt.Println("looking for", card_to_find, "card is at index", GetItem(favorite_cards[:], card_to_find))
	fmt.Println("looking for", card_to_find, "card is at index", GetItem(cards, card_to_find))
	// Replace
	index := 2
	new_card := 6
	fmt.Println("cards:", cards)
	cards = SetItem(cards, index, new_card)
	fmt.Println("replaced card at index 2", cards)
	// Append
	new_card = 9
	index = 100
	cards = SetItem(cards, index, new_card)
	fmt.Println("added card to index 100 (end)", cards)
	new_card = 2
	index = -1
	cards = SetItem(cards, index, new_card)
	fmt.Println("added card to index -1 (end)", cards)
	// Prepend
	cards = []int{3, 2, 6, 4, 8}
	fmt.Println("cards:", cards)
	cards = PrependItems(cards, 5, 1)
	fmt.Println("added cards to beginning of deck", cards)
	// Remove
	cards = []int{3, 2, 6, 4, 8}
	fmt.Println("cards:", cards)
	cards = RemoveItem(cards, 11)
	cards = RemoveItem(cards, 4)
	fmt.Println("Removed 4", cards)
}
