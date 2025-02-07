package main

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layout := "2006-01-02 15:04:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("time parse err:", err)
	}
	return t
}

func HasPassed() {

}

func IsAfternoonAppointment() {

}

func Description() {

}

func AnniversaryDate() {

}

func TimeTest() {
	date := "7/25/2019 13:45:00"
	fmt.Println(fmt.Sprintf("date %s, scheduled for %s", date, Schedule(date)))
	fmt.Println("TODO")
}
