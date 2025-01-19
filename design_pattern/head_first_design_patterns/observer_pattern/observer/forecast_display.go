package observer

import (
	"fmt"
	"tucker_algorithm/head_first_design_patterns/observer_pattern/domain"
)

type forecastDisplay struct {
	currentPressure float32
	lastPressure    float32
	weatherData     domain.Subject
}

func InitForecastDisplay(weatherData domain.Subject) domain.Observer {
	o := &forecastDisplay{
		currentPressure: 29.2,
		weatherData:     weatherData,
	}
	weatherData.RegisterObserver(o)
	return o
}

func (d *forecastDisplay) Display() {
	fmt.Print("Forecast: ")
	if d.currentPressure > d.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if d.currentPressure == d.lastPressure {
		fmt.Println("More of the same")
	} else if d.currentPressure < d.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}

func (d *forecastDisplay) Update() {
	d.lastPressure = d.currentPressure
	d.currentPressure = d.weatherData.Pressure()

	d.Display()
}
