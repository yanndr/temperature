//Package temperature provides Temperature struct and a convert method to deal with different temperature scale.
package temperature

import (
	"errors"
	"fmt"
	"math"
)

// ErrNilInputTemperature is an error when the input temperature is nil.
var ErrNilInputTemperature = errors.New("Input temperature can't be nil")

// Temperature is an Temperature interface.
type Temperature interface {
	Unit() Unit
	SetUnit(Unit)
	Value() float64
	SetValue(float64)
	String() string
	SetTemperature(Temperature)
}

type temperature struct {
	value float64
	unit  Unit
}

//NewTemperature returns a temperature.
func NewTemperature(value float64, unit Unit) Temperature {
	return &temperature{value, unit}
}

func (t *temperature) String() string {
	return fmt.Sprintf("%v %s", round(t.Value(), 2), t.unit.Text)
}

// Unit retruns the unit of the temperature.
func (t *temperature) Unit() Unit {
	return t.unit
}

// SetUnit set the unit of the temperature and update the temperature.
func (t *temperature) SetUnit(u Unit) {
	t.value = u.FromKelvin(t.Unit().ToKelvin(t.Value()))
	t.unit = u
}

// Value returns the value of the temperature.
func (t *temperature) Value() float64 {
	return t.value
}

// SetValue set the value of the temperature.
func (t *temperature) SetValue(v float64) {
	t.value = v
}

// SetTemperature set the temperature value from any other unit temperature.
func (t *temperature) SetTemperature(temp Temperature) {
	t.value = t.unit.FromKelvin(temp.Unit().ToKelvin(temp.Value()))
}

func round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}

// Convert a temperature to different unit.
func Convert(input Temperature, unit Unit) (Temperature, error) {
	if input == nil {
		return nil, ErrNilInputTemperature
	}

	val := unit.FromKelvin(input.Unit().ToKelvin(input.Value()))
	return &temperature{val, unit}, nil
}

//Equals returns true if two temperature are equals.
func Equals(a Temperature, b Temperature) bool {
	if a == b {
		return true
	}

	if a.Unit().Text == b.Unit().Text {
		return a.Value() == b.Value()
	}

	ka := a.Unit().ToKelvin(a.Value())
	kb := b.Unit().ToKelvin(b.Value())
	return round(ka, 2) == round(kb, 2)
}

type handlerTemperatureChanged func(Temperature)

type monitoredTemperature struct {
	Temperature
	handler handlerTemperatureChanged
}

//NewMonitoredTemperature returns a thermometer that send the updated temperature to a channel.
func NewMonitoredTemperature(unit Unit, handler handlerTemperatureChanged) Temperature {
	return &monitoredTemperature{
		Temperature: NewTemperature(0, unit),
		handler:     handler,
	}
}

// SetTemperature set the temperature value from any other unit temperature.
func (t *monitoredTemperature) SetTemperature(temp Temperature) {

	t.Temperature.SetTemperature(temp)

	if t.handler != nil {
		t.handler(t)
	}
}

// SetValue set the value of the temperature.
func (t *monitoredTemperature) SetValue(v float64) {

	t.Temperature.SetValue(v)

	if t.handler != nil {
		t.handler(t)
	}
}
