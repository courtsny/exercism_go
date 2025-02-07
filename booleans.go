package main

import "fmt"

// a fast attack can be made if the knight is sleeping
func CanFastAttack(isKnightAwake bool) bool {
	return !isKnightAwake
}

// The group can be spied upon if at least one of them is awake.
func CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake bool) bool {
	return isKnightAwake || isArcherAwake || isPrisonerAwake
}

// The prisoner can be signaled using bird sounds
// if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(isArcherAwake, isPrisonerAwake bool) bool {
	return !isArcherAwake && isPrisonerAwake
}

/*
If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep.
The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.

If Annalyn does not have her dog then she and the prisoner must be very sneaky!
Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but

if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.
*/
func CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent bool) bool {
	if isDogPresent {
		if !isArcherAwake {
			return true
		}
	} else {
		if isPrisonerAwake && !isArcherAwake && !isKnightAwake {
			return true
		}
	}
	return false
}

func BooleansTest() {
	isKnightAwake := true
	fmt.Println("can fast attack?", CanFastAttack(isKnightAwake))
	isKnightAwake = false
	isArcherAwake := true
	isPrisonerAwake := false
	fmt.Println("can spy?", CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake))
	isArcherAwake = false
	isPrisonerAwake = true
	fmt.Println("can signal?", CanSignalPrisoner(isArcherAwake, isPrisonerAwake))
	isKnightAwake = false
	isArcherAwake = true
	isPrisonerAwake = false
	isDogPresent := false
	fmt.Println("can free?", CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent))
}
