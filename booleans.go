package main

import "fmt"

// a fast attack can be made if the knight is sleeping
func CanFastAttack(isKnightAwake bool) bool {
	fmt.Printf("isKnightAwake is %t, so CanFastAttack is %t\n", isKnightAwake, !isKnightAwake)
	return !isKnightAwake
}

// The group can be spied upon if at least one of them is awake.
func CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake bool) bool {
	canSpy := isKnightAwake || isArcherAwake || isPrisonerAwake
	fmt.Printf("isKnightAwake is %t, isArcherAwake is %t, and isPrisonerAwake is %t, so CanSpy is %t\n",
		isKnightAwake, isArcherAwake, isPrisonerAwake, canSpy)
	return canSpy
}

// The prisoner can be signaled using bird sounds
// if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(isArcherAwake, isPrisonerAwake bool) bool {
	canSignalPrisoner := !isArcherAwake && isPrisonerAwake
	fmt.Printf("isArcherAwake is %t and isPrisonerAwake is %t, so CanSignalPrisoner is %t\n",
		isArcherAwake, isPrisonerAwake, canSignalPrisoner)
	return canSignalPrisoner
}

/*
If Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep.
The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.

If Annalyn does not have her dog then she and the prisoner must be very sneaky!
Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping, but

if the prisoner is sleeping they can't be rescued: the prisoner would be startled by Annalyn's sudden appearance and wake up the knight and archer.
*/
func CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent bool) bool {
	fmt.Printf("isKnightAwake is %t, isArcherAwake is %t, isPrisonerAwake is %t, and isDogPresent is %t, ",
		isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent)
	if isDogPresent {
		if isPrisonerAwake && !isArcherAwake {
			fmt.Printf("so canFreePrisoner is true\n")
			return true
		}
	} else {
		if isPrisonerAwake && !isArcherAwake && !isKnightAwake {
			fmt.Printf("so canFreePrisoner is true\n")
			return true
		}
	}
	fmt.Printf("so canFreePrisoner is false\n")
	return false
}

func BooleansTest() {
	isKnightAwake := true
	CanFastAttack(isKnightAwake)
	isKnightAwake = false
	isArcherAwake := true
	isPrisonerAwake := false
	CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake)
	isArcherAwake = false
	isPrisonerAwake = true
	CanSignalPrisoner(isArcherAwake, isPrisonerAwake)
	isKnightAwake = false
	isArcherAwake = true
	isPrisonerAwake = false
	isDogPresent := false
	CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent)
}
