package main

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table_number int, seatmate_name, table_direction string, distance_to_table float64) string {
	var return_string = Welcome(name)
	return_string += " " + fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table_number, table_direction, distance_to_table)
	return_string += " " + fmt.Sprintf("You will be sitting next to %s.", seatmate_name)
	return return_string
}

func PackagesTest() {
	name := "Person"
	fmt.Println(Welcome(name))
	age := 50
	fmt.Println(HappyBirthday(name, age))
	name = "Christiane"
	table := 27
	seatmate_name := "Frank"
	direction := "on the left"
	distance := 23.7834298
	fmt.Println(AssignTable(name, table, seatmate_name, direction, distance))
}
