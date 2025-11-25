// Package weather tells you today's weather.
package weather

var (
// CurrentCondition tells you hows the weather.
	CurrentCondition string
// CurrentLocation tells you where you are.
	CurrentLocation  string
)
/* Forecast is a function that takes the city's name and the current weather condition, and returns today's weather forecast. */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
