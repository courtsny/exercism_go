package main

import "fmt"

func NeedsLicense(kind string) bool {
	if kind == "bike" {
		return false
	}
	return true
}

func ChooseVehicle(car1, car2 string) string {
	if car2 < car1 {
		return car2
	}
	return car1
}

/*
If the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new.
If it is at least 10 years old, it costs 50%.
If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.
*/
func CalculateResellPrice(originalPrice int, age int) int {
	var cost_percentage = 0.7
	if age < 3 {
		cost_percentage = 0.8
	}
	if age > 10 {
		cost_percentage = 0.5
	}
	return int(cost_percentage * float64(originalPrice))
}

func IfTest() {
	kind := "bike"
	fmt.Printf("a %s needs a license? %t\n", kind, NeedsLicense(kind))
	kind = "car"
	fmt.Printf("a %s needs a license? %t\n", kind, NeedsLicense(kind))
	car1 := "Volkswagen Beetle"
	car2 := "Toyota Corolla"
	fmt.Printf("between %s and %s, you should get a %s\n", car1, car2, ChooseVehicle(car1, car2))
	fmt.Printf("between %s and %s, you should get a %s\n", car2, car1, ChooseVehicle(car2, car1))
	car_age := 1
	og_car_price := 1000
	fmt.Printf("my %d year old car cost $%d and is now worth $%d\n", car_age, og_car_price, CalculateResellPrice(og_car_price, car_age))
	car_age = 5
	fmt.Printf("my %d year old car cost $%d and is now worth $%d\n", car_age, og_car_price, CalculateResellPrice(og_car_price, car_age))
	car_age = 15
	fmt.Printf("my %d year old car cost $%d and is now worth $%d\n", car_age, og_car_price, CalculateResellPrice(og_car_price, car_age))
}
