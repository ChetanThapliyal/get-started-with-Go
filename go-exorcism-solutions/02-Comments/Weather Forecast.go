// Package weather provides the latest weather conditions
//of user in the city he/she is.
package weather

//CurrentCondition represents current weather conditions.
var CurrentCondition string
//CurrentLocation represents the current location/city of user.
var CurrentLocation string

//Forecast describes user about the weather in his/her current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
