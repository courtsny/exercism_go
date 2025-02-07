package main

import "fmt"

var layers = []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}

func PreparationTime(ingredients []string, mins int) int {
	// Default to 2 mins
	if mins == 0 {
		mins = 2
	}
	return len(ingredients) * mins
}

// For each noodle layer in your lasagna, you will need 50 grams of noodles.
// For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
func Quantities(ing []string) (int, float64) {
	var noodles_in_grams = 0
	var sauce_in_liters = 0.0
	for _, element := range ing {
		if element == "noodles" {
			noodles_in_grams += 50
		} else if element == "sauce" {
			sauce_in_liters += 0.2
		}
	}
	return noodles_in_grams, sauce_in_liters
}

func AddSecretIngredient(friend_list, my_list []string) {
	var my_last_index = len(my_list) - 1
	var friend_last_index = len(friend_list) - 1
	my_list[my_last_index] = friend_list[friend_last_index]
}

func FunctionsTest() {
	fmt.Println(fmt.Sprintf("it'll take %d mins to cook %d layers", PreparationTime(layers, 3), len(layers)))
	fmt.Println(fmt.Sprintf("it'll take %d mins to cook %d layers", PreparationTime(layers, 0), len(layers)))
	var noodles, sauce = Quantities(layers)
	fmt.Println(fmt.Sprintf("i'll need %d grams of noodles and %.2f liters of sauce", noodles, sauce))
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	fmt.Println("my list:", myList)
	fmt.Println("friend's list:", friendsList)
	AddSecretIngredient(friendsList, myList)
	fmt.Println("my list after secret ingredient:", myList)
}
