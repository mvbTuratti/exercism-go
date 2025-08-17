// Package weather provides tools to forecast the weather in Goblinocus.
package weather

// CurrentCondition hold the string representation of the current weather condition.
var CurrentCondition string

// CurrentLocation holds the string representation of the current location that shall be forecast.
var CurrentLocation string

// Forecast returns the string message of location and their respective forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
