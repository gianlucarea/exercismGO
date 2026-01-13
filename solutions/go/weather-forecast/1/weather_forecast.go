// Package weather provides tools for describing current weather conditions.
package weather

var (
	// CurrentCondition stores the current weather condition (e.g. sunny, rainy).
	CurrentCondition string

	// CurrentLocation stores the name of the current location.
	CurrentLocation string
)

// Forecast sets the current location and condition and returns a formatted
// weather report string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

