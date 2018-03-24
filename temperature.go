//Package temperature provides Temperature struct and a convert method to deal with different temperature scale.
package temperature

import (
	"errors"
	"fmt"
	"math"
)

// ErrNilArgument is an error when the argument is nil.
var ErrNilArgument = errors.New("argument can't be nil")

type temperatureChangeHandler func(Temperature)

//Stringer provides String method.
type Stringer interface {
	String() string
}

// Temperature provides all the function needed for a temperature.
type Temperature interface {
	Stringer
	Value() float64
	SetValue(float64)
	Unit() Convertible
	SetUnit(Convertible)

	SetTemperature(Temperature)
	SetTemperateChangeHandler(temperatureChangeHandler)
}

type temperature struct {
	v       float64
	unit    Convertible
	handler temperatureChangeHandler
}

//New returns a new Temperature.
func New(v float64, unit Convertible) Temperature {
	return &temperature{v: v, unit: unit}
}

//NewWithHandler returns a new temperature with a handler to handle temperature changes.
func NewWithHandler(v float64, unit Convertible, handler temperatureChangeHandler) Temperature {
	return &temperature{v: v, unit: unit, handler: handler}
}

func (t temperature) String() string {
	return fmt.Sprintf("%v %s", round(float64(t.v), 2), t.unit)
}

// Value returns the value of the temperature.
func (t temperature) Value() float64 {
	return t.v
}

// // SetValue set the value of the temperature.
func (t temperature) SetValue(v float64) {
	t.v = v

	if t.handler != nil {
		t.handler(t)
	}
}

func (t temperature) Unit() Convertible {
	return t.unit
}

// SetUnit set the unit of the temperature and update the temperature.
func (t temperature) SetUnit(u Convertible) {
	t.v = u.FromKelvin(t.unit.ToKelvin(t.v)).Value()
	t.unit = u
}

// SetTemperature set the temperature value from any other unit temperature.
func (t temperature) SetTemperature(temp Temperature) {
	val := t.unit.FromKelvin(temp.Unit().ToKelvin(temp.Value())).Value()
	t.SetValue(val)
}

// SetTemperateChangeHandler is a setter for the temperature change handler.
func (t temperature) SetTemperateChangeHandler(h temperatureChangeHandler) {
	t.handler = h
}

func round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}

// Convert a temperature to different unit.
func Convert(input Temperature, output Convertible) (Temperature, error) {
	if input == nil || output == nil {
		return nil, ErrNilArgument
	}

	return output.FromKelvin(input.Unit().ToKelvin(input.Value())), nil
}

//Equals returns true if two temperature are equals.
func Equals(a Temperature, b Temperature) bool {
	if a == b {
		return true
	}

	if a.Unit() == b.Unit() {
		return a.Value() == b.Value()
	}

	ka := a.Unit().ToKelvin(a.Value())
	kb := b.Unit().ToKelvin(b.Value())
	return round(ka, 2) == round(kb, 2)
}
