package main

import (
	"fmt"
	"math/rand"
)

var animals = [8]string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

func RollADie() int {
	return (rand.Intn(19) + 1)
}

func GenerateWandEnergy() float64 {
	// A random number between 0 - 11 + A random float between 0 and 1 (excluding 0 and 1)
	return float64(rand.Intn(12)) + rand.Float64()
}

func ShuffleAnimals() []string {
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals[:]
}

func RandomnessTest() {
	fmt.Printf("rolling a random int: %d\n", RollADie())
	fmt.Printf("generate a random floating point: %.2f\n", GenerateWandEnergy())
	fmt.Printf("animals before shuffling: %s\n", animals)
	ShuffleAnimals()
	fmt.Printf("animals after shuffling: %s\n", animals)
}
