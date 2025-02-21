package main

import (
	"fmt"
)

var birdsPerDay = []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

func TotalBirdCount(birds []int) int {
	var bird_count = 0
	for _, element := range birds {
		bird_count += element
	}
	return bird_count
}

func BirdsInWeek(birds []int, week int) int {
	// week 1 == 0, w2 = 6, w3 = 13
	var start_index = (week - 1) * 7
	var end_index = start_index + 7

	if end_index > len(birds) {
		fmt.Println("no data for week", week)
		return -1
	}
	//fmt.Println(birds[start_index:end_index])
	return TotalBirdCount(birds[start_index:end_index])
}

func FixBirdCountLog(birds []int) []int {
	for index, _ := range birds {
		if index%2 == 0 {
			birds[index] += 1
		}
	}
	return birds
}

func ForTest() {
	fmt.Printf("i saw %d birds in %d days\n", TotalBirdCount(birdsPerDay), len(birdsPerDay))
	week := 2
	fmt.Printf("in week %d i saw %d birds\n", week, BirdsInWeek(birdsPerDay, week))
	week = 1
	fmt.Printf("in week %d i saw %d birds\n", week, BirdsInWeek(birdsPerDay, week))
	fmt.Println("old log:", birdsPerDay)
	var new_log = FixBirdCountLog(birdsPerDay)
	fmt.Println("new log:", new_log)
	fmt.Printf("after fixing the log, i actually saw %d birds in %d days\n", TotalBirdCount(new_log), len(birdsPerDay))
}
