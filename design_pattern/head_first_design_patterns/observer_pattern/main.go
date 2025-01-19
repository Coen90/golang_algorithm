package main

import (
	"tucker_algorithm/head_first_design_patterns/observer_pattern/observer"
	"tucker_algorithm/head_first_design_patterns/observer_pattern/subject"
)

func main() {
	weatherData := subject.InitWeatherData()

	observer.InitCurrentConditionsDisplay(weatherData)
	observer.InitStatisticsDisplay(weatherData)
	observer.InitForecastDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
