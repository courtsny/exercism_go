package main

import "fmt"

func CalculateWorkingCarsPerHour(carsProducedPerHour, successRate int) float64 {
	carsProducedPerHour_Float := float64(carsProducedPerHour)
	successRate_float := float64(successRate) / 100.0
	return carsProducedPerHour_Float * successRate_float
}

func CalculateWorkingCarsPerMinute(carsProducedPerHour, successRate int) int {
	carsPerHour := CalculateWorkingCarsPerHour(carsProducedPerHour, successRate)
	return int(carsPerHour / 60)
}

func CalculateCost(numberOfCars int) uint {
	var groupsOfTen = numberOfCars / 10
	var theRest = numberOfCars % 10
	return uint(groupsOfTen)*95000 + uint(theRest)*10000
}

func ArithmeticTest() {
	carCount := 1547
	carRate := 90
	fmt.Println("cars per hour =", carCount, "success rate =", carRate)
	fmt.Println("working cars per hour =", CalculateWorkingCarsPerHour(carCount, carRate))
	carCount = 1105
	carRate = 90
	fmt.Println("cars per hour =", carCount, "success rate =", carRate)
	fmt.Println("working cars per minute =", CalculateWorkingCarsPerMinute(carCount, carRate))
	carCount = 37
	fmt.Println("cars to make =", carCount)
	fmt.Println("cost =", CalculateCost(carCount))
	carCount = 21
	fmt.Println("cars to make =", carCount)
	fmt.Println("cost =", CalculateCost(carCount))
}
