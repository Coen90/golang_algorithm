package observer

import (
	"fmt"
	"tucker_algorithm/head_first_design_patterns/observer_pattern/domain"
)

type statisticsDisplay struct {
	maxTemp     float32
	minTemp     float32
	tempSum     float32
	numReadings int
	weatherData domain.Subject
}

func InitStatisticsDisplay(weatherData domain.Subject) domain.Observer {
	o := &statisticsDisplay{
		maxTemp:     0.0,
		minTemp:     200,
		tempSum:     0.0,
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(o)
	return o
}

func (d *statisticsDisplay) Display() {
	fmt.Printf("Avg/Max/Min temperature = %.1f / %.1f / %.1f\n",
		d.tempSum/float32(d.numReadings), d.maxTemp, d.minTemp)
}

func (d *statisticsDisplay) Update() {
	d.tempSum += d.weatherData.Temperature()
	d.numReadings++
	if d.weatherData.Temperature() > d.maxTemp {
		d.maxTemp = d.weatherData.Temperature()
	}
	if d.weatherData.Temperature() < d.minTemp {
		d.minTemp = d.weatherData.Temperature()
	}

	d.Display()
}
