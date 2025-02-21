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
		fmt.Printf("card %d DNE in %v\n", card_to_remove, cards)
	}
	return cards
}

func SlicesTest() {
	// Favorite
	fmt.Println("favorite cards are:", favorite_cards)
	// Search
	card_to_find := 2
	cards := []int{3, 7, 8, 1, 4}
	fmt.Printf("looking for card %d in %v. card is at index %d\n",
		card_to_find, favorite_cards, GetItem(favorite_cards[:], card_to_find))
	fmt.Printf("looking for card %d in %v. card is at index %d\n",
		card_to_find, cards, GetItem(cards, card_to_find))
	// Replace
	index := 2
	new_card := 6
	fmt.Println("cards:", cards)
	old_card := cards[index]
	cards = SetItem(cards, index, new_card)
	fmt.Printf("replaced card %d at index %d with %d\n", old_card, index, new_card)
	// Append
	new_card = 9
	index = 100
	fmt.Println("cards:", cards)
	cards = SetItem(cards, index, new_card)
	fmt.Printf("added card %d to index %d (end)\n", new_card, index)
	new_card = 2
	index = -1
	fmt.Println("cards:", cards)
	cards = SetItem(cards, index, new_card)
	fmt.Printf("added card %d to index -1 (end)\n", new_card)
	// Prepend
	fmt.Println("cards:", cards)
	cards = PrependItems(cards, 5, 1)
	fmt.Printf("added cards to beginning of deck\n")
	// Remove
	fmt.Println("cards:", cards)
	cardToRemove := 11
	cards = RemoveItem(cards, cardToRemove) // should throw an error since DNE
	fmt.Println("cards:", cards)
	cardToRemove = 4
	cards = RemoveItem(cards, cardToRemove)
	fmt.Printf("Removed %d: %v\n", cardToRemove, cards)
}
