package main

import (
	"fmt"
)

var layers = []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}

func PreparationTime(ingredients []string, mins int) int {
	// Default to 2 mins
	if mins == 0 {
		mins = 2
	}
	return len(ingredients) * mins
}

// For each noodle layer, need 50 grams of noodles.
// For each sauce layer, need 0.2 liters of sauce.
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
	fmt.Printf("it'll take %d mins to cook %d layers\n", PreparationTime(layers, 3), len(layers))
	fmt.Printf("it'll take %d mins to cook %d layers\n", PreparationTime(layers, 0), len(layers))
	var noodles, sauce = Quantities(layers)
	fmt.Printf("i'll need %d grams of noodles and %.2f liters of sauce\n", noodles, sauce)
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	fmt.Println("my list:", myList)
	fmt.Println("friend's list:", friendsList)
	AddSecretIngredient(friendsList, myList)
	fmt.Println("my list after secret ingredient:", myList)
}

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Day1Records(r Record) bool {
	return r.Day == 1
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

func ByCategory(category string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == category {
			return true
		}
		return false
	}
}

func TotalByPeriod(r []Record, p DaysPeriod) float64 {
	sum := 0.0
	filteredRecords := Filter(r, ByDaysPeriod(p))
	for _, x := range filteredRecords {
		sum += x.Amount
	}
	return sum
}

func Filter(records []Record, recordFunc func(Record) bool) []Record {
	var filteredRecords = []Record{}
	for _, r := range records {
		if recordFunc(r) == true {
			filteredRecords = append(filteredRecords, r)
		}
	}
	return filteredRecords
}

type UnknownCategoryError struct {
	category string
}

func (u UnknownCategoryError) Error() string {
	return fmt.Sprintf("unkown category %s", u.category)
}

func CategoryExpenses(r []Record, p DaysPeriod, category string) (float64, error) {
	filteredRecords := Filter(r, ByCategory(category))
	if len(filteredRecords) < 1 {
		return 0.0, UnknownCategoryError{category: category}
	}
	expenses := TotalByPeriod(filteredRecords, p)
	return expenses, nil
}

func FirstClassTest() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}
	day1 := Day1Records
	fmt.Printf("filtering %v where day = 1: %v\n", records, Filter(records, day1))

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	period := DaysPeriod{From: 1, To: 15}
	fmt.Printf("filtering %v where day is between %v: %v\n", records, period, Filter(records, ByDaysPeriod(period)))

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	category := "groceries"
	fmt.Printf("filtering %v where category = %s: %v\n", records, category, Filter(records, ByCategory(category)))

	records = []Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}

	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	fmt.Printf("summing totals in %v where day is between %v: %v\n", records, p1, TotalByPeriod(records, p1))
	fmt.Printf("summing totals in %v where day is between %v: %v\n", records, p2, TotalByPeriod(records, p2))

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	category = "entertainment"
	expenses, err := CategoryExpenses(records, p1, category)
	fmt.Printf("category expenses for category %s with days between %v: %.1f with error '%v'\n", category, p1, expenses, err)

	category = "rent"
	expenses, err = CategoryExpenses(records, p1, category)
	fmt.Printf("category expenses for category %s with days between %v: %.1f with error '%v'\n", category, p1, expenses, err)

	expenses, err = CategoryExpenses(records, p2, category)
	fmt.Printf("category expenses for category %s with days between %v: %.1f with error '%v'\n", category, p2, expenses, err)
}
