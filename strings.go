package main

import (
	"fmt"
	"strings"
)

func WelcomeMessage(name string) {
	fmt.Println("Welcome to the Tech Palace,", strings.ToUpper(name))
}

func AddBorder(message string, star_count int) string {
	var stars = strings.Repeat("*", star_count)
	return stars + string('\n') + message + string('\n') + stars
}

func CleanUpMessage(message string) string {
	var no_stars = strings.ReplaceAll(message, "*", "")
	return no_stars
}

func StringsTest() {
	WelcomeMessage("Courtney")
	starMessage := AddBorder("Welcome!", 10)
	fmt.Println("Adding stars:\n", starMessage)
	fmt.Println("Removing stars:", CleanUpMessage(starMessage))
}
