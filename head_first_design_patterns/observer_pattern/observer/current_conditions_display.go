package observer

import (
	"fmt"
	"tucker_algorithm/head_first_design_patterns/observer_pattern/domain"
)

type currentConditionsDisplay struct {
	temperature float32
	humidity    float32
	weatherData domain.Subject
}

func InitCurrentConditionsDisplay(weatherData domain.Subject) domain.Observer {
	c := &currentConditionsDisplay{
		weatherData: weatherData,
	}
	weatherData.RegisterObserver(c)
	return c
}

func (d *currentConditionsDisplay) Display() {
	fmt.Printf("현재 상태: 온도 %.1fF, 습도 %.1f\n", d.temperature, d.humidity)
}

func (d *currentConditionsDisplay) Update() {
	d.temperature = d.weatherData.Temperature()
	d.humidity = d.weatherData.Humidity()
	d.Display()
}
