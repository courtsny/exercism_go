package main

import (
	"errors"
	"fmt"
)

var amount_err = errors.New("food amount error")
var factor_err = errors.New("fattening factor error")

type FodderCalculator struct {
	amount           float64
	amountErr        error
	triggerAmountErr bool
	factor           float64
	factorErr        error
	triggerFactorErr bool
}

type InvalidCowsError struct {
	num_cows int
	message  string
}

func (i InvalidCowsError) Error() string {
	return fmt.Sprintf("[number of cows] cows are invalid: %d cows triggered '%s'", i.num_cows, i.message)
}

func (fc FodderCalculator) FodderAmount(amount int) (float64, error) {
	if fc.triggerAmountErr {
		return 0, amount_err
	}
	return fc.amount, fc.amountErr
}

func (fc FodderCalculator) FatteningFactor() (float64, error) {
	if fc.triggerFactorErr {
		return 0, factor_err
	}
	return fc.factor, fc.factorErr
}

func DivideFood(fc FodderCalculator, num_cows int) (float64, error) {
	amount, err := fc.FodderAmount(num_cows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	div_value := (amount * factor) / float64(num_cows)
	return div_value, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, num_cows int) (float64, error) {
	err := ValidateNumberOfCows(num_cows)
	if err != nil {
		return 0, err
	}
	return DivideFood(fc, num_cows)
}

func ValidateNumberOfCows(num_cows int) error {
	if num_cows < 0 {
		return InvalidCowsError{num_cows: num_cows, message: "there are no negative cows"}
	} else if num_cows == 0 {
		return InvalidCowsError{num_cows: num_cows, message: "no cows don't need to eat"}
	}
	return nil
}

func ErrorsTest() {
	fc := FodderCalculator{amount: 50, factor: 1.5}
	food, err := DivideFood(fc, 5)
	fmt.Printf("divided food %.0f with error {%v}\n", food, err) // happy path :)

	fc.triggerAmountErr = true
	food, err = DivideFood(fc, 5)
	fmt.Printf("divided food %.0f with error {%v}\n", food, err)

	fc.triggerAmountErr = false
	fc.triggerFactorErr = true
	food, err = DivideFood(fc, 5)
	fmt.Printf("divided food %.0f with error {%v}\n", food, err)
	fc.triggerFactorErr = false // remove all errors

	food, err = ValidateInputAndDivideFood(fc, 5)
	fmt.Printf("validated & divided food %.0f with error {%v}\n", food, err)
	food, err = ValidateInputAndDivideFood(fc, 0)
	fmt.Printf("validated & divided food %.0f with error {%v}\n", food, err)
	food, err = ValidateInputAndDivideFood(fc, -2)
	fmt.Printf("validated & divided food %.0f with error {%v}\n", food, err)
}
