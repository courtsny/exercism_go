package main

import "fmt"

type Car struct {
	battery       int
	battery_drain int
	speed         int
	distance      int
}

func NewCar(speed_in_meters, battery_drain_percentage int) Car {
	return Car{
		battery:       100,
		battery_drain: battery_drain_percentage,
		speed:         speed_in_meters,
		distance:      0,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance_in_meters int) Track {
	return Track{
		distance: distance_in_meters,
	}
}

func Drive(car Car) Car {
	car.battery -= car.battery_drain
	car.distance += car.speed
	return car
}

func CanFinish(car Car, track Track) bool {
	// How many laps will it take to finish the track?
	var lap_count = track.distance / car.speed

	// Can the car make it that many laps?
	var battery_drainage = lap_count * car.battery_drain
	if battery_drainage < 100 {
		fmt.Printf("it'll take %d laps and battery will be at %d\n", lap_count, car.battery-battery_drainage)
		return true
	}
	/*
		var laps = 0
		for i := car.battery - car.battery_drain; i > 0; i -= car.battery_drain {
			car.battery = i
			car.distance += car.speed
			laps += 1
			fmt.Println(fmt.Sprintf("car distance = %d, car battery = %d, lap count = %d", car.distance, car.battery, laps))
			if car.distance >= track.distance {
				return true
			}
		}
	*/
	return false
}

// Pointer allows the car's values to be changed in the function
func (car *Car) Drive() {
	car.battery -= car.battery_drain
	car.distance += car.speed
}

// Non-pointer allows "read only" of car's values
func (car Car) DisplayDistance() {
	fmt.Printf("Driven %d meters\n", car.distance)
}

func (car Car) DisplayBattery() {
	fmt.Printf("Battery at %d%%\n", car.battery)
}

func (car Car) CanFinish(track_distance int) bool {
	// If less than the maximum distance, yes
	if car.battery_drain*car.speed < track_distance {
		return true
	}
	return false
}

func StructsTest() {
	car := NewCar(5, 2)
	fmt.Println("new car with 5 speed and battery drain 2:", car)
	fmt.Println("new track with 800m distance", NewTrack(800))
	car = Drive(car)
	fmt.Println("drove car one lap", car)
	car = NewCar(5, 2)
	track := NewTrack(100)
	fmt.Printf("can the car %v make it around the track %v? %t\n", car, track, CanFinish(car, track))

	fmt.Printf("\n *** using built in methods on the car now ***\n\n")
	new_car := NewCar(5, 2)
	fmt.Printf("new car battery = %d and distance = %d\n", new_car.battery, new_car.distance)
	new_car.Drive()
	new_car.DisplayDistance()
	new_car.DisplayBattery()
	fmt.Printf("can the car %v make it around the track %v? %t\n", new_car, track, new_car.CanFinish(track.distance))
}
