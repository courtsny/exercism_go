package main

import "fmt"

type Stringer interface {
	String() string
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (md MeteorologyData) String() string {
	return fmt.Sprintf("In %s the temperature is %s. Wind %s at %s with %d%% humidity.", md.location, md.temperature.String(),
		md.windDirection, md.windSpeed.String(), md.humidity)
}

func StringersTest() {
	cu := Celsius
	fu := Fahrenheit
	fmt.Printf("cu is %s, fu is %s\n", cu.String(), fu.String())

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	fmt.Printf("celsius stringer: %s, fahrenheit stringer: %s\n", celsiusTemp.String(), fahrenheitTemp.String())

	mph := MilesPerHour
	kph := KmPerHour
	fmt.Printf("mph is %s, kph is %s\n", mph.String(), kph.String())

	windSpeedNow := Speed{
		magnitude: 18,
		unit:      KmPerHour,
	}

	windSpeedYesterday := Speed{
		magnitude: 22,
		unit:      MilesPerHour,
	}

	fmt.Printf("kmh speed is: %s\nmph speed is: %s\n", windSpeedNow.String(), windSpeedYesterday.String())

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}

	fmt.Printf("meteorology data is: %s\n", sfData)
}
