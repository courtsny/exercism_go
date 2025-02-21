package main

import "fmt"

func InterestRate(balance float32) float32 {
	// default rate for > 0
	interest_rate := float32(0.03213)

	// switch would've been nice here but not supported for non-int types
	if balance >= 0 {
		interest_rate = float32(0.005)
	} else if balance >= 1000 {
		interest_rate = float32(0.01621)
	} else if balance >= 5000 {
		interest_rate = float32(0.02475)
	}
	return interest_rate
}

func Interest(balance float32) float64 {
	return float64(balance * InterestRate(balance))
}

func AnnualBalanceUpdate(balance float32) float64 {
	return float64(balance) + Interest(balance)
}

func YearsBeforeDesiredBalance(balance, targetBalance float32) int {
	currentBalance := balance
	for i := 0; ; i++ {
		if currentBalance >= targetBalance {
			return i
		}
		currentBalance += float32(Interest(currentBalance))
	}
}

func FloatTest() {
	balance := float32(200.75)
	fmt.Printf("if i have %.2f then my interest rate is %.5f\n", balance, InterestRate(balance))
	fmt.Printf("if i have %.2f then my interest after 1 year is %.6f\n", balance, Interest(balance))
	fmt.Printf("if i have %.2f then my balance after 1 year is %.5f\n", balance, AnnualBalanceUpdate(balance))
	desiredBalance := float32(214.88)
	fmt.Printf("if i have %.2f and want to have %.2f, it will take me %d years\n",
		balance, desiredBalance, YearsBeforeDesiredBalance(balance, desiredBalance))
}
