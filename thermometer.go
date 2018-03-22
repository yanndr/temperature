package temperature

type monitoredTemperature struct {
	Temperature
	updateChan chan<- interface{}
}

//NewMonitoredTemperature returns a thermometer that send the updated temperature to a channel.
func NewMonitoredTemperature(unit Unit, ch chan<- interface{}) Temperature {
	return &monitoredTemperature{
		Temperature: NewTemperature(0, unit),
		updateChan:  ch,
	}
}

// SetTemperature set the temperature value from any other unit temperature.
func (t *monitoredTemperature) SetTemperature(temp Temperature) {

	t.Temperature.SetTemperature(temp)

	if t.updateChan != nil {
		t.updateChan <- t
	}
}

// SetValue set the value of the temperature.
func (t *monitoredTemperature) SetValue(v float64) {

	t.Temperature.SetValue(v)

	if t.updateChan != nil {
		t.updateChan <- t
	}
}
