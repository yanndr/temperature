//Package temperature provides Temperature struct to deal with different temperature scale.
package temperature

import (
	"errors"
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
	Rankine    struct{ Temperature }
	Delisle    struct{ Temperature }
	Reaumur    struct{ Temperature }
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

func NewRankine(value float64) Rankine {
	return Rankine{
		Temperature: Temperature{Value: value},
	}
}

func NewDelisle(value float64) Delisle {
	return Delisle{
		Temperature: Temperature{Value: value},
	}
}

func NewReaumur(value float64) Reaumur {
	return Reaumur{
		Temperature: Temperature{Value: value},
	}
}

//Round returns a round number of a float64.
func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	val := int(num*output + math.Copysign(0.5, num))
	return float64(val) / output
}

var ErrNilOutputTemperature = errors.New("Output temperature can't be nil")
var ErrNilInputTemperature = errors.New("Input temperature can't be nil")

// Convert a temperature to different unit. It updated the output temperature value.
func Convert(input ToKelvinConvertible, output FromKelvinConvertible) error {
	if output == nil {
		return ErrNilOutputTemperature
	}

	if input == nil {
		return ErrNilInputTemperature
	}

	output.FromKelvin(input.ToKelvin())
	return nil
}
