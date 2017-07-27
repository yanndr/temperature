//Package temperature provides Temperature struct to deal with different temperature scale.
package temperature

import (
	"math"
)

type ToKelvinConvertible interface {
	ToKelvin() Kelvin
}

type FromKelvinConvertible interface {
	FromKelvin(Valuer)
}

type Valuer interface {
	GetValue() float64
}

type Temperature struct {
	Value float64
}

// GetValue returns the value of the temperature.
func (t Temperature) GetValue() float64 {
	return t.Value
}

type (
	Celsius    struct{ Temperature }
	Kelvin     struct{ Temperature }
	Fahrenheit struct{ Temperature }
)

func NewCelsius(value float64) Celsius {
	return Celsius{
		Temperature: Temperature{Value: value},
	}
}

func NewKelvin(value float64) Kelvin {
	return Kelvin{
		Temperature: Temperature{Value: value},
	}
}

func NewFahrenheit(value float64) Fahrenheit {
	return Fahrenheit{
		Temperature: Temperature{Value: value},
	}
}

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}

// Convert a temperature to different unit. It updated the output temperature value.
func Convert(input ToKelvinConvertible, output FromKelvinConvertible) {
	output.FromKelvin(input.ToKelvin())
}
