package main

import "fmt"

// / Basics
const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func CookTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return CookTime(numberOfLayers) + actualMinutesInOven
}

func BasicsTest() {
	layers := 3
	minutes := 20
	fmt.Println("number of layers =", layers, "actual minutes in oven =", minutes)
	fmt.Println(ElapsedTime(layers, minutes))
}
