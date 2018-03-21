//Package temperature provides Temperature struct and a convert method to deal with different temperature scale.
package temperature

import (
	"errors"
	"fmt"
	"math"
)

// Unit represents a temperature unit
type Unit struct {
	Text       string
	ToKelvin   func(float64) float64
	FromKelvin func(float64) float64
}

// ErrNilOutputTemperature is an error when the output temperature is nil.
var ErrNilOutputTemperature = errors.New("Output temperature can't be nil")

// ErrNilInputTemperature is an error when the input temperature is nil.
var ErrNilInputTemperature = errors.New("Input temperature can't be nil")

var (
	// Celsius temperature unit.
	Celsius = Unit{
		Text:       "°C",
		ToKelvin:   func(v float64) float64 { return v + 273.15 },
		FromKelvin: func(v float64) float64 { return v - 273.15 },
	}
	//Kelvin temperature unit.
	Kelvin = Unit{
		Text:       "K",
		ToKelvin:   func(v float64) float64 { return v },
		FromKelvin: func(v float64) float64 { return v },
	}
	// Fahrenheit temperature unit.
	Fahrenheit = Unit{
		Text:       "°F",
		ToKelvin:   func(v float64) float64 { return (v + 459.67) * 5 / 9 },
		FromKelvin: func(v float64) float64 { return v*9/5 - 459.67 },
	}
	// Rankine temperature unit.
	Rankine = Unit{
		Text:       "°Ra",
		ToKelvin:   func(v float64) float64 { return v * 5 / 9 },
		FromKelvin: func(v float64) float64 { return v * 9 / 5 },
	}
	// Delisle temperature unit.
	Delisle = Unit{
		Text:       "°D",
		ToKelvin:   func(v float64) float64 { return 373.15 - v*2/3 },
		FromKelvin: func(v float64) float64 { return (373.15 - v) * 3 / 2 },
	}
	// Reaumur temperature unit.
	Reaumur = Unit{
		Text:       "°Re",
		ToKelvin:   func(v float64) float64 { return v*5/4 + 273.15 },
		FromKelvin: func(v float64) float64 { return (v - 273.15) * 4 / 5 },
	}
)

// Temperature is an Temperature interface.
type Temperature interface {
	Unit() Unit
	Value() float64
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

// Value returns the value of the temperature.
func (t *temperature) Value() float64 {
	return t.value
}

// SetTemperature set the temperature value from any other unit temperature.
func (t *temperature) SetTemperature(temp Temperature) {
	t.value = t.unit.FromKelvin(temp.Unit().ToKelvin(temp.Value()))
}

// NewCelsius returns a Celsius temperature initialized with the value.
func NewCelsius(value float64) Temperature {
	return &temperature{value: value, unit: Celsius}
}

// NewKelvin returns a Kelvin temperature initialized with the value.
func NewKelvin(value float64) Temperature {
	return &temperature{value: value, unit: Kelvin}
}

// NewFahrenheit returns a Fahrenheit temperature initialized with the value.
func NewFahrenheit(value float64) Temperature {
	return &temperature{value: value, unit: Fahrenheit}
}

// NewRankine returns a Rankine Temperature nitialized with the value.
func NewRankine(value float64) Temperature {
	return &temperature{value: value, unit: Rankine}
}

// NewDelisle returns a Delisle temperature initialized with the value.
func NewDelisle(value float64) Temperature {
	return &temperature{value: value, unit: Delisle}
}

// NewReaumur returns a Reaumur temperature initialized with the value.
func NewReaumur(value float64) Temperature {
	return &temperature{value: value, unit: Reaumur}
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
