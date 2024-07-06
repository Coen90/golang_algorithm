package subject

import "tucker_algorithm/head_first_design_patterns/observer_pattern/domain"

type weatherData struct {
	observer    []domain.Observer
	temperature float32
	humidity    float32
	pressure    float32
}

func InitWeatherData() domain.Subject {
	return &weatherData{
		observer: make([]domain.Observer, 0),
	}
}

func (w *weatherData) RegisterObserver(observer domain.Observer) {
	if w.getRegisterObserverIndex(observer) != -1 {
		return
	}
	w.observer = append(w.observer, observer)
}

func (w *weatherData) RemoveObserver(observer domain.Observer) {
	idx := w.getRegisterObserverIndex(observer)
	if idx == -1 {
		return
	}
	if idx == 0 {
		w.observer = w.observer[1:]
		return
	}
	w.observer = append(w.observer[:idx], w.observer[idx+1:]...)
}

func (w *weatherData) SetMeasurements(temperature, humidity, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}

func (w *weatherData) Temperature() float32 {
	return w.temperature
}

func (w *weatherData) Humidity() float32 {
	return w.humidity
}

func (w *weatherData) Pressure() float32 {
	return w.pressure
}

func (w *weatherData) notifyObservers() {
	for i := range w.observer {
		w.observer[i].Update()
	}
}

func (w *weatherData) getRegisterObserverIndex(observer domain.Observer) int {
	result := -1
	for i := range w.observer {
		if w.observer[i] == observer {
			result = i
		}
	}
	return result
}

func (w *weatherData) measurementsChanged() {
	w.notifyObservers()
}
