package domain

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	SetMeasurements(temperature, humidity, pressure float32)
	Temperature() float32
	Humidity() float32
	Pressure() float32
}
