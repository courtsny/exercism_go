package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseTime(date string) time.Time {
	// The layout example has to use this specific timestamp info
	// Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("parse time failed with err:", err)
	}
	return t
}

func HasPassed(date string) bool {
	formatted := ParseTime(date)
	if formatted.Compare(time.Now()) < 0 {
		return true
	}
	return false
}

func IsAfternoonAppointment(date string) bool {
	formatted := ParseTime(date)
	if formatted.Hour() >= 12 && formatted.Hour() < 18 {
		return true
	}
	return false
}

func Description(date string) string {
	f := ParseTime(date)
	var sb strings.Builder
	sb.WriteString(f.Weekday().String() + ", ")
	sb.WriteString(f.Month().String() + " " + strconv.Itoa(f.Day()) + ", " + strconv.Itoa(f.Year()))
	sb.WriteString(" at " + strconv.Itoa(f.Hour()) + ":" + strconv.Itoa(f.Minute()) + ".")
	return sb.String()
}

func AnniversaryDate() time.Time {
	layout := "1/02/2006 15:04:05"
	anniversary := "9/15/2025 00:00:00"
	t, err := time.Parse(layout, anniversary)
	if err != nil {
		fmt.Println("anniversary date time parse failed with err:", err)
	}
	return t
}

func TimeTest() {
	date := "7/25/2019 13:45:00"
	fmt.Printf("date %s, scheduled for %s\n", date, ParseTime(date))
	fmt.Printf("appointment on %s has passed: %t\n", date, HasPassed(date))
	fmt.Printf("appointment on %s is in the afternoon: %t\n", date, IsAfternoonAppointment(date))
	fmt.Printf("description of appointment is: %s\n", Description(date))
	fmt.Printf("the anniversary date this year is %s\n", AnniversaryDate().String())
}
